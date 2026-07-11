#!/usr/bin/env python3
"""Validate the generated agent-facing Hugo outputs."""

from __future__ import annotations

import re
import sys
import xml.etree.ElementTree as ET
from html.parser import HTMLParser
from pathlib import Path
from urllib.parse import unquote, urljoin, urlparse


CANONICAL_ORIGIN = "https://www.qrlecosystem.com"
CANONICAL_HOST = "www.qrlecosystem.com"
LLMS_LINK_PATTERN = re.compile(
    r"^- \[([^\]]+)\]\((https://[^)\s]+)\)(?:: .+)?$"
)
FORBIDDEN_ORIGINS = (
    "https://qrlecosystem.io",
    "http://qrlecosystem.io",
    "https://qrlecosystem.com",
    "http://qrlecosystem.com",
    "http://www.qrlecosystem.com",
)


class PageMetadataParser(HTMLParser):
    def __init__(self) -> None:
        super().__init__()
        self.canonical_links: list[str] = []
        self.markdown_links: list[str] = []

    def handle_starttag(
        self, tag: str, attrs: list[tuple[str, str | None]]
    ) -> None:
        if tag != "link":
            return

        values = dict(attrs)
        relationships = (values.get("rel") or "").split()
        href = values.get("href")
        if not href:
            return
        if "canonical" in relationships:
            self.canonical_links.append(href)
        if "alternate" in relationships and values.get("type") == "text/markdown":
            self.markdown_links.append(href)


def published_path(publish_dir: Path, url_path: str) -> Path:
    return publish_dir / unquote(url_path).lstrip("/")


def main() -> int:
    publish_dir = Path(sys.argv[1]) if len(sys.argv) > 1 else Path("website/public")
    errors: list[str] = []

    required_files = {
        "llms.txt",
        "robots.txt",
        "sitemap.xml",
        "index.json",
        "index.html.md",
    }
    for relative_path in sorted(required_files):
        if not (publish_dir / relative_path).is_file():
            errors.append(f"Missing required output: {relative_path}")

    sitemap_path = publish_dir / "sitemap.xml"
    sitemap_urls: list[str] = []
    if sitemap_path.is_file():
        try:
            sitemap_root = ET.parse(sitemap_path).getroot()
            sitemap_urls = [
                element.text.strip()
                for element in sitemap_root.findall("{*}url/{*}loc")
                if element.text
            ]
        except (ET.ParseError, OSError) as error:
            errors.append(f"Unable to parse sitemap.xml: {error}")

    if not sitemap_urls:
        errors.append("sitemap.xml contains no page URLs")

    expected_markdown_files: set[Path] = set()
    for page_url in sitemap_urls:
        parsed = urlparse(page_url)
        if parsed.scheme != "https" or parsed.netloc != CANONICAL_HOST:
            errors.append(f"Non-canonical sitemap URL: {page_url}")
            continue
        if not parsed.path.endswith("/"):
            errors.append(f"Sitemap content URL must end with '/': {page_url}")
            continue
        if "/projects/active/" in parsed.path or "/projects/archived/" in parsed.path:
            errors.append(f"Project sitemap URL exposes status directory: {page_url}")
        if any(marker in parsed.path for marker in (".md", "llms.txt", "robots.txt")):
            errors.append(f"Agent output must not appear in sitemap.xml: {page_url}")

        output_dir = published_path(publish_dir, parsed.path)
        html_path = output_dir / "index.html"
        markdown_path = output_dir / "index.html.md"
        expected_markdown_files.add(markdown_path)

        if not html_path.is_file():
            errors.append(f"Sitemap URL has no generated HTML file: {page_url}")
            continue
        if not markdown_path.is_file():
            errors.append(f"Sitemap URL has no Markdown alternative: {page_url}")

        html = html_path.read_text(encoding="utf-8")
        parser = PageMetadataParser()
        parser.feed(html)

        if parser.canonical_links != [page_url]:
            errors.append(
                f"{html_path}: expected one canonical link to {page_url}, "
                f"found {parser.canonical_links}"
            )

        expected_markdown_url = urljoin(page_url, "index.html.md")
        if parser.markdown_links != [expected_markdown_url]:
            errors.append(
                f"{html_path}: expected one Markdown alternate link to "
                f"{expected_markdown_url}, found {parser.markdown_links}"
            )

    actual_markdown_files = set(publish_dir.rglob("index.html.md"))
    missing_markdown = expected_markdown_files - actual_markdown_files
    extra_markdown = actual_markdown_files - expected_markdown_files
    for path in sorted(missing_markdown):
        errors.append(f"Missing Markdown output: {path}")
    for path in sorted(extra_markdown):
        errors.append(f"Markdown output is not represented in sitemap.xml: {path}")

    forbidden_markup = ("<!doctype", "<html", "<head", "<body", "<script")
    for markdown_path in sorted(actual_markdown_files):
        markdown = markdown_path.read_text(encoding="utf-8")
        if not markdown.startswith("# "):
            errors.append(f"{markdown_path}: Markdown output must start with an H1")
        if markdown.startswith("---"):
            errors.append(f"{markdown_path}: Markdown output contains front matter")
        lowered = markdown.lower()
        for marker in forbidden_markup:
            if marker in lowered:
                errors.append(f"{markdown_path}: contains HTML shell marker {marker!r}")

    llms_path = publish_dir / "llms.txt"
    if llms_path.is_file():
        llms_text = llms_path.read_text(encoding="utf-8")
        llms_lines = llms_text.splitlines()
        if not llms_lines or llms_lines[0] != "# QRL Ecosystem Index":
            errors.append("llms.txt must start with the site H1")

        headings = [line for line in llms_lines if line.startswith("## ")]
        expected_headings = ["## Core Resources", "## Projects", "## Optional"]
        if headings != expected_headings:
            errors.append(
                f"llms.txt sections must be {expected_headings}; found {headings}"
            )

        first_section = next(
            (index for index, line in enumerate(llms_lines) if line.startswith("## ")),
            len(llms_lines),
        )
        if not any(line.startswith("> ") for line in llms_lines[1:first_section]):
            errors.append("llms.txt must include a summary blockquote before its sections")

        links: list[tuple[str, str]] = []
        for line_number, line in enumerate(llms_lines, start=1):
            if not line.startswith("- ["):
                continue
            match = LLMS_LINK_PATTERN.fullmatch(line)
            if not match:
                errors.append(f"llms.txt:{line_number}: invalid file-list entry")
                continue
            links.append((match.group(1), match.group(2)))

        urls = [url for _, url in links]
        if len(urls) != len(set(urls)):
            errors.append("llms.txt contains duplicate URLs")

        for _, url in links:
            parsed = urlparse(url)
            if parsed.scheme != "https":
                errors.append(f"llms.txt link is not HTTPS: {url}")
            if parsed.netloc == CANONICAL_HOST:
                if parsed.query or parsed.fragment:
                    errors.append(f"Internal llms.txt link must be canonical: {url}")
                    continue
                if (
                    "/projects/active/" in parsed.path
                    or "/projects/archived/" in parsed.path
                ):
                    errors.append(f"llms.txt project link exposes status directory: {url}")
                local_path = published_path(publish_dir, parsed.path)
                if not local_path.is_file():
                    errors.append(f"Internal llms.txt link does not resolve: {url}")

        try:
            projects_start = llms_lines.index("## Projects") + 1
            optional_start = llms_lines.index("## Optional")
            project_names = [
                match.group(1)
                for line in llms_lines[projects_start:optional_start]
                if (match := LLMS_LINK_PATTERN.fullmatch(line))
            ]
            if project_names != sorted(project_names, key=str.casefold):
                errors.append("llms.txt project links are not alphabetically ordered")
        except ValueError:
            pass

    robots_path = publish_dir / "robots.txt"
    if robots_path.is_file():
        robots_lines = [
            line.strip()
            for line in robots_path.read_text(encoding="utf-8").splitlines()
            if line.strip()
        ]
        expected_robots = [
            "User-agent: *",
            "Allow: /",
            f"Sitemap: {CANONICAL_ORIGIN}/sitemap.xml",
        ]
        if robots_lines != expected_robots:
            errors.append(
                f"robots.txt must contain {expected_robots}; found {robots_lines}"
            )

    checked_suffixes = {".html", ".md", ".txt", ".xml", ".json"}
    for output_path in publish_dir.rglob("*"):
        if not output_path.is_file() or output_path.suffix not in checked_suffixes:
            continue
        text = output_path.read_text(encoding="utf-8")
        for origin in FORBIDDEN_ORIGINS:
            if origin in text:
                errors.append(f"{output_path}: contains forbidden origin {origin}")

    if errors:
        print("AGENT OUTPUT VALIDATION ERRORS:")
        for error in errors:
            print(f"  - {error}")
        return 1

    print(
        f"Agent outputs are valid: {len(sitemap_urls)} HTML pages, "
        f"{len(actual_markdown_files)} Markdown alternatives."
    )
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
