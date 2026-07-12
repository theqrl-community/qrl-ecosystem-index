(() => {
    const galleries = Array.from(document.querySelectorAll("[data-project-gallery]"));

    galleries.forEach((gallery) => {
        const slides = Array.from(gallery.querySelectorAll("[data-gallery-slide]"));
        const thumbnails = Array.from(gallery.querySelectorAll("[data-gallery-thumbnail]"));
        const caption = gallery.querySelector("[data-gallery-caption]");
        const position = gallery.querySelector("[data-gallery-position]");
        const dialog = gallery.querySelector("[data-gallery-dialog]");
        const dialogImage = gallery.querySelector("[data-gallery-dialog-image]");
        const dialogCaption = gallery.querySelector("[data-gallery-dialog-caption]");
        const dialogPosition = gallery.querySelector("[data-gallery-dialog-position]");
        const closeButton = gallery.querySelector("[data-gallery-close]");
        let currentIndex = 0;
        let lastOpener = null;

        if (!slides.length) {
            return;
        }

        const selectScreenshot = (requestedIndex) => {
            currentIndex = (requestedIndex + slides.length) % slides.length;
            const activeSlide = slides[currentIndex];
            const activeImage = activeSlide.querySelector("img");
            const activeCaption = activeSlide.dataset.caption || activeImage.alt;

            slides.forEach((slide, index) => {
                slide.hidden = index !== currentIndex;
            });
            thumbnails.forEach((thumbnail, index) => {
                if (index === currentIndex) {
                    thumbnail.setAttribute("aria-current", "true");
                } else {
                    thumbnail.removeAttribute("aria-current");
                }
            });

            if (caption) {
                caption.textContent = activeCaption;
            }
            if (position) {
                position.textContent = String(currentIndex + 1);
            }
            if (dialogImage) {
                dialogImage.src = activeImage.currentSrc || activeImage.src;
                dialogImage.alt = activeCaption;
            }
            if (dialogCaption) {
                dialogCaption.textContent = activeCaption;
            }
            if (dialogPosition) {
                dialogPosition.textContent = String(currentIndex + 1);
            }
        };

        gallery.querySelectorAll("[data-gallery-direction]").forEach((button) => {
            button.addEventListener("click", () => {
                selectScreenshot(currentIndex + Number(button.dataset.galleryDirection));
            });
        });

        thumbnails.forEach((thumbnail) => {
            thumbnail.addEventListener("click", () => {
                selectScreenshot(Number(thumbnail.dataset.galleryIndex));
            });
        });

        gallery.querySelectorAll("[data-gallery-open]").forEach((opener) => {
            opener.addEventListener("click", (event) => {
                if (!dialog || typeof dialog.showModal !== "function") {
                    return;
                }
                event.preventDefault();
                lastOpener = opener;
                dialog.showModal();
                document.body.classList.add("gallery-dialog-open");
                if (closeButton) {
                    closeButton.focus();
                }
            });
        });

        if (dialog && closeButton) {
            closeButton.addEventListener("click", () => dialog.close());
            dialog.addEventListener("close", () => {
                document.body.classList.remove("gallery-dialog-open");
                if (lastOpener) {
                    lastOpener.focus();
                }
            });
        }

        gallery.addEventListener("keydown", (event) => {
            if (event.key === "Escape" && dialog && dialog.open) {
                event.preventDefault();
                dialog.close();
            } else if (event.key === "ArrowLeft") {
                event.preventDefault();
                selectScreenshot(currentIndex - 1);
            } else if (event.key === "ArrowRight") {
                event.preventDefault();
                selectScreenshot(currentIndex + 1);
            }
        });

        const enableSwipe = (surface) => {
            if (!surface) {
                return;
            }
            let startX = null;
            let startY = null;

            surface.addEventListener("pointerdown", (event) => {
                if (event.isPrimary === false) {
                    return;
                }
                startX = event.clientX;
                startY = event.clientY;
            });
            surface.addEventListener("pointerup", (event) => {
                if (startX === null || startY === null) {
                    return;
                }
                const distanceX = event.clientX - startX;
                const distanceY = event.clientY - startY;
                startX = null;
                startY = null;

                if (Math.abs(distanceX) < 50 || Math.abs(distanceX) <= Math.abs(distanceY)) {
                    return;
                }
                selectScreenshot(currentIndex + (distanceX < 0 ? 1 : -1));
            });
            surface.addEventListener("pointercancel", () => {
                startX = null;
                startY = null;
            });
        };

        enableSwipe(gallery.querySelector("[data-gallery-swipe]"));
        enableSwipe(gallery.querySelector("[data-gallery-dialog-swipe]"));
        selectScreenshot(0);
        gallery.classList.add("project-gallery-ready");
    });
})();
