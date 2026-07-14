import tempfile
import unittest
from pathlib import Path

from validate_gallery import MAX_SCREENSHOT_BYTES, validate_gallery


class ValidateGalleryTest(unittest.TestCase):
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

    def validate(self, gallery):
        return validate_gallery(
            self.yaml_path,
            {"id": "example-project", "gallery": gallery},
            self.root,
        )

    def test_optional_gallery_is_valid(self):
        self.assertEqual(
            validate_gallery(
                self.yaml_path, {"id": "example-project"}, self.root
            ),
            [],
        )

    def test_one_to_ten_ordered_mixed_items_are_valid(self):
        gallery = []
        for index in range(10):
            if index % 2 == 0:
                filename = f"screen-{index}.webp"
                self.add_image(filename)
                gallery.append(
                    {
                        "type": "image",
                        "path": f"example-project/{filename}",
                        "caption": f"Screenshot {index}",
                    }
                )
            else:
                gallery.append(
                    {
                        "type": "youtube",
                        "id": f"Video{index:06d}",
                        "caption": f"Video {index}",
                    }
                )
        self.assertEqual(self.validate(gallery), [])

    def test_invalid_submissions_are_rejected(self):
        self.add_image("valid.png")
        self.add_image("duplicate.png")
        self.add_image("oversized.webp", MAX_SCREENSHOT_BYTES + 1)

        invalid_cases = {
            "external image URL": [
                {
                    "type": "image",
                    "path": "https://example.com/screen.png",
                    "caption": "External",
                }
            ],
            "youtube URL instead of ID": [
                {
                    "type": "youtube",
                    "id": "https://youtu.be/M7lc1UVf-VE",
                    "caption": "External",
                }
            ],
            "malformed youtube ID": [
                {"type": "youtube", "id": "too-short", "caption": "Malformed"}
            ],
            "duplicate youtube ID": [
                {"type": "youtube", "id": "M7lc1UVf-VE", "caption": "First"},
                {"type": "youtube", "id": "M7lc1UVf-VE", "caption": "Second"},
            ],
            "youtube path conflict": [
                {
                    "type": "youtube",
                    "id": "M7lc1UVf-VE",
                    "path": "example-project/valid.png",
                    "caption": "Conflict",
                }
            ],
            "image ID conflict": [
                {
                    "type": "image",
                    "path": "example-project/valid.png",
                    "id": "M7lc1UVf-VE",
                    "caption": "Conflict",
                }
            ],
            "unknown type": [
                {"type": "video", "id": "M7lc1UVf-VE", "caption": "Unknown"}
            ],
            "missing type": [
                {"path": "example-project/valid.png", "caption": "Missing type"}
            ],
            "traversal": [
                {
                    "type": "image",
                    "path": "example-project/../screen.png",
                    "caption": "Traversal",
                }
            ],
            "absolute path": [
                {
                    "type": "image",
                    "path": "/example-project/screen.png",
                    "caption": "Absolute",
                }
            ],
            "wrong project": [
                {
                    "type": "image",
                    "path": "another-project/screen.png",
                    "caption": "Wrong project",
                }
            ],
            "missing file": [
                {
                    "type": "image",
                    "path": "example-project/missing.png",
                    "caption": "Missing",
                }
            ],
            "duplicate": [
                {
                    "type": "image",
                    "path": "example-project/duplicate.png",
                    "caption": "First",
                },
                {
                    "type": "image",
                    "path": "example-project/duplicate.png",
                    "caption": "Second",
                },
            ],
            "unsupported format": [
                {
                    "type": "image",
                    "path": "example-project/screen.svg",
                    "caption": "Unsupported",
                }
            ],
            "missing caption": [
                {"type": "image", "path": "example-project/valid.png"}
            ],
            "long caption": [
                {
                    "type": "image",
                    "path": "example-project/valid.png",
                    "caption": "x" * 161,
                }
            ],
            "oversized": [
                {
                    "type": "image",
                    "path": "example-project/oversized.webp",
                    "caption": "Too large",
                }
            ],
            "empty gallery": [],
            "eleven gallery items": [
                {
                    "type": "youtube",
                    "id": f"Video{index:06d}",
                    "caption": str(index),
                }
                for index in range(11)
            ],
        }

        for name, gallery in invalid_cases.items():
            with self.subTest(name=name):
                self.assertTrue(self.validate(gallery))


if __name__ == "__main__":
    unittest.main()
