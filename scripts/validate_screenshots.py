#!/usr/bin/env python3
"""Validate project screenshot metadata and local asset files."""

from __future__ import annotations

import sys
from pathlib import Path, PurePosixPath
from typing import Any

import yaml


MAX_SCREENSHOTS = 10
MAX_SCREENSHOT_BYTES = 2 * 1024 * 1024
MAX_CAPTION_LENGTH = 160
SUPPORTED_EXTENSIONS = {".png", ".jpg", ".jpeg", ".webp"}


def validate_screenshots(
    yaml_path: Path, data: dict[str, Any], repository_root: Path
) -> list[str]:
    screenshots = data.get("screenshots")
    if screenshots is None:
        return []

    errors: list[str] = []
    if not isinstance(screenshots, list):
        return [f"{yaml_path}: screenshots must be a list"]
    if not 1 <= len(screenshots) <= MAX_SCREENSHOTS:
        errors.append(
            f"{yaml_path}: screenshots must contain between 1 and {MAX_SCREENSHOTS} entries"
        )

    project_id = data.get("id")
    seen_paths: set[str] = set()

    for index, screenshot in enumerate(screenshots):
        label = f"screenshots[{index}]"
        if not isinstance(screenshot, dict):
            errors.append(f"{yaml_path}: {label} must be an object")
            continue

        caption = screenshot.get("caption")
        if not isinstance(caption, str) or not caption.strip():
            errors.append(f"{yaml_path}: {label}.caption is required")
        elif len(caption) > MAX_CAPTION_LENGTH:
            errors.append(
                f"{yaml_path}: {label}.caption exceeds {MAX_CAPTION_LENGTH} characters"
            )
        elif "\n" in caption or "\r" in caption:
            errors.append(f"{yaml_path}: {label}.caption must be a single line")

        screenshot_path = screenshot.get("path")
        if not isinstance(screenshot_path, str) or not screenshot_path:
            errors.append(f"{yaml_path}: {label}.path is required")
            continue
        if screenshot_path in seen_paths:
            errors.append(f"{yaml_path}: {label}.path is duplicated: {screenshot_path}")
            continue
        seen_paths.add(screenshot_path)

        if screenshot_path.startswith(("http://", "https://", "//")):
            errors.append(
                f"{yaml_path}: {label}.path must be local, not a URL: {screenshot_path}"
            )
            continue
        if "\\" in screenshot_path:
            errors.append(
                f"{yaml_path}: {label}.path must use forward slashes: {screenshot_path}"
            )
            continue

        relative_path = PurePosixPath(screenshot_path)
        if relative_path.is_absolute() or any(
            part in {"", ".", ".."} for part in relative_path.parts
        ):
            errors.append(
                f"{yaml_path}: {label}.path must not be absolute or traverse directories: {screenshot_path}"
            )
            continue
        if len(relative_path.parts) != 2 or relative_path.parts[0] != project_id:
            errors.append(
                f"{yaml_path}: {label}.path must be inside '{project_id}/': {screenshot_path}"
            )
            continue
        if relative_path.suffix not in SUPPORTED_EXTENSIONS:
            errors.append(
                f"{yaml_path}: {label}.path must use a lowercase PNG, JPEG, or WebP extension: {screenshot_path}"
            )
            continue

        full_path = repository_root / "images" / "screenshots" / Path(*relative_path.parts)
        if not full_path.is_file():
            errors.append(
                f"{yaml_path}: {label} file not found: {screenshot_path} (expected at {full_path})"
            )
            continue
        if full_path.stat().st_size > MAX_SCREENSHOT_BYTES:
            errors.append(
                f"{yaml_path}: {label} exceeds the 2 MB file-size limit: {screenshot_path}"
            )

    return errors


def project_yaml_files(repository_root: Path):
    return (
        yaml_file
        for yaml_file in (repository_root / "projects").rglob("*.yaml")
        if yaml_file.name != "template.yaml"
    )


def main() -> int:
    repository_root = Path(__file__).resolve().parent.parent
    errors: list[str] = []

    for yaml_file in project_yaml_files(repository_root):
        with yaml_file.open() as file:
            data = yaml.safe_load(file) or {}
        errors.extend(validate_screenshots(yaml_file, data, repository_root))

    if errors:
        print("SCREENSHOT VALIDATION ERRORS:")
        print("\n".join(f"  - {error}" for error in errors))
        return 1

    print("All project screenshots are valid.")
    return 0


if __name__ == "__main__":
    sys.exit(main())
