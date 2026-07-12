import tempfile
import unittest
from pathlib import Path

from validate_screenshots import MAX_SCREENSHOT_BYTES, validate_screenshots


class ValidateScreenshotsTest(unittest.TestCase):
    def setUp(self):
        self.temporary_directory = tempfile.TemporaryDirectory()
        self.root = Path(self.temporary_directory.name)
        self.yaml_path = self.root / "projects" / "active" / "example-project.yaml"
        self.yaml_path.parent.mkdir(parents=True)

    def tearDown(self):
        self.temporary_directory.cleanup()

    def add_image(self, filename: str, size: int = 8) -> None:
        path = self.root / "images" / "screenshots" / "example-project" / filename
        path.parent.mkdir(parents=True, exist_ok=True)
        path.write_bytes(b"x" * size)

    def validate(self, screenshots):
        return validate_screenshots(
            self.yaml_path,
            {"id": "example-project", "screenshots": screenshots},
            self.root,
        )

    def test_optional_screenshots_are_valid(self):
        self.assertEqual(
            validate_screenshots(
                self.yaml_path, {"id": "example-project"}, self.root
            ),
            [],
        )

    def test_one_to_ten_ordered_screenshots_are_valid(self):
        screenshots = []
        for index in range(10):
            filename = f"screen-{index}.webp"
            self.add_image(filename)
            screenshots.append(
                {
                    "path": f"example-project/{filename}",
                    "caption": f"Screenshot {index}",
                }
            )
        self.assertEqual(self.validate(screenshots), [])

    def test_invalid_submissions_are_rejected(self):
        self.add_image("valid.png")
        self.add_image("duplicate.png")
        self.add_image("oversized.webp", MAX_SCREENSHOT_BYTES + 1)

        invalid_cases = {
            "external URL": [
                {"path": "https://example.com/screen.png", "caption": "External"}
            ],
            "traversal": [
                {"path": "example-project/../screen.png", "caption": "Traversal"}
            ],
            "absolute path": [
                {"path": "/example-project/screen.png", "caption": "Absolute"}
            ],
            "wrong project": [
                {"path": "another-project/screen.png", "caption": "Wrong project"}
            ],
            "missing file": [
                {"path": "example-project/missing.png", "caption": "Missing"}
            ],
            "duplicate": [
                {"path": "example-project/duplicate.png", "caption": "First"},
                {"path": "example-project/duplicate.png", "caption": "Second"},
            ],
            "unsupported format": [
                {"path": "example-project/screen.svg", "caption": "Unsupported"}
            ],
            "missing caption": [
                {"path": "example-project/valid.png"}
            ],
            "long caption": [
                {"path": "example-project/valid.png", "caption": "x" * 161}
            ],
            "oversized": [
                {"path": "example-project/oversized.webp", "caption": "Too large"}
            ],
            "eleven screenshots": [
                {"path": "example-project/valid.png", "caption": str(index)}
                for index in range(11)
            ],
        }

        for name, screenshots in invalid_cases.items():
            with self.subTest(name=name):
                self.assertTrue(self.validate(screenshots))


if __name__ == "__main__":
    unittest.main()
