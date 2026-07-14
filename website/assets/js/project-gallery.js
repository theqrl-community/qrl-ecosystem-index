(() => {
    const youtubeThumbnailVariants = ["maxresdefault", "hqdefault"];
    const youtubeThumbnailStates = new WeakMap();

    const youtubeThumbnailURL = (videoID, variant) => `https://i.ytimg.com/vi/${videoID}/${variant}.jpg`;

    const isHighResolutionThumbnail = (image) => image.naturalWidth >= 640 && image.naturalHeight >= 360;

    const markThumbnailMissing = (image) => {
        image.hidden = true;
        image.closest(".project-gallery-video-poster, .project-gallery-thumbnail-video")?.classList.add("project-gallery-video-thumbnail-missing");
    };

    const markThumbnailAvailable = (image) => {
        image.hidden = false;
        image.closest(".project-gallery-video-poster, .project-gallery-thumbnail-video")?.classList.remove("project-gallery-video-thumbnail-missing");
    };

    const loadYouTubeThumbnail = (image, videoID) => {
        if (!videoID) {
            markThumbnailMissing(image);
            return;
        }

        const previousState = youtubeThumbnailStates.get(image);
        const state = {
            request: (previousState?.request || 0) + 1,
            variantIndex: 0,
        };
        youtubeThumbnailStates.set(image, state);
        image.dataset.galleryYoutubeThumbnailVideoId = videoID;

        const isCurrentRequest = () => youtubeThumbnailStates.get(image) === state;
        const tryNextVariant = () => {
            if (!isCurrentRequest()) {
                return;
            }

            const variant = youtubeThumbnailVariants[state.variantIndex];
            if (!variant) {
                markThumbnailMissing(image);
                return;
            }

            image.hidden = false;
            image.onload = () => {
                if (!isCurrentRequest()) {
                    return;
                }
                if (variant === "maxresdefault" && !isHighResolutionThumbnail(image)) {
                    state.variantIndex += 1;
                    tryNextVariant();
                    return;
                }
                image.dataset.galleryYoutubeThumbnailVariant = variant;
                markThumbnailAvailable(image);
            };
            image.onerror = () => {
                if (!isCurrentRequest()) {
                    return;
                }
                state.variantIndex += 1;
                tryNextVariant();
            };
            image.src = youtubeThumbnailURL(videoID, variant);
        };

        tryNextVariant();
    };

    const galleries = Array.from(document.querySelectorAll("[data-project-gallery]"));

    galleries.forEach((gallery) => {
        const slides = Array.from(gallery.querySelectorAll("[data-gallery-slide]"));
        const thumbnails = Array.from(gallery.querySelectorAll("[data-gallery-thumbnail]"));
        const caption = gallery.querySelector("[data-gallery-caption]");
        const position = gallery.querySelector("[data-gallery-position]");
        const dialog = gallery.querySelector("[data-gallery-dialog]");
        const dialogImage = gallery.querySelector("[data-gallery-dialog-image]");
        const dialogVideo = gallery.querySelector("[data-gallery-dialog-video]");
        const dialogVideoPoster = gallery.querySelector("[data-gallery-dialog-video-open]");
        const dialogVideoThumbnail = dialogVideoPoster?.querySelector("[data-gallery-youtube-thumbnail]");
        const dialogVideoPlayer = gallery.querySelector("[data-gallery-dialog-video-player]");
        const dialogCaption = gallery.querySelector("[data-gallery-dialog-caption]");
        const dialogPosition = gallery.querySelector("[data-gallery-dialog-position]");
        const closeButton = gallery.querySelector("[data-gallery-close]");
        let currentIndex = 0;
        let lastOpener = null;

        if (!slides.length) {
            return;
        }

        const wrappedIndex = (requestedIndex, length) => (requestedIndex + length) % length;

        const createVideoIframe = (videoID, title) => {
            const iframe = document.createElement("iframe");
            iframe.src = `https://www.youtube-nocookie.com/embed/${videoID}?autoplay=1&playsinline=1`;
            iframe.title = title || "YouTube video";
            iframe.loading = "lazy";
            iframe.referrerPolicy = "strict-origin-when-cross-origin";
            iframe.allow = "accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share";
            iframe.allowFullscreen = true;
            return iframe;
        };

        const resetVideo = (slide) => {
            if (slide.dataset.galleryMediaType !== "youtube") {
                return;
            }
            const poster = slide.querySelector("[data-gallery-video-open]");
            const player = slide.querySelector("[data-gallery-video-player]");
            if (player) {
                player.replaceChildren();
                player.hidden = true;
            }
            if (poster) {
                poster.hidden = false;
            }
        };

        const resetDialogVideo = () => {
            if (dialogVideoPlayer) {
                dialogVideoPlayer.replaceChildren();
                dialogVideoPlayer.hidden = true;
            }
            if (dialogVideoPoster) {
                dialogVideoPoster.hidden = false;
            }
        };

        const syncDialog = (slide) => {
            if (!dialog) {
                return;
            }
            const slideIndex = slides.indexOf(slide);
            if (slideIndex < 0) {
                return;
            }

            const isVideo = slide.dataset.galleryMediaType === "youtube";
            const activeImage = slide.querySelector("[data-gallery-image]");
            const activeCaption = slide.dataset.caption || activeImage?.alt || "Project media";
            resetDialogVideo();
            dialog.classList.toggle("project-gallery-dialog-active-video", isVideo);

            if (dialogImage) {
                dialogImage.hidden = isVideo;
                if (!isVideo) {
                    dialogImage.src = activeImage?.currentSrc || activeImage?.src || "";
                    dialogImage.alt = activeCaption;
                }
            }

            if (dialogVideo) {
                dialogVideo.hidden = !isVideo;
                if (isVideo) {
                    const videoID = slide.dataset.youtubeId;
                    dialogVideo.dataset.youtubeId = videoID || "";
                    if (dialogVideoPoster && videoID) {
                        dialogVideoPoster.href = `https://www.youtube.com/watch?v=${videoID}`;
                        dialogVideoPoster.setAttribute("aria-label", `Play video: ${activeCaption}`);
                    }
                    if (dialogVideoThumbnail && videoID) {
                        if (dialogVideoThumbnail.dataset.galleryYoutubeThumbnailVideoId !== videoID) {
                            loadYouTubeThumbnail(dialogVideoThumbnail, videoID);
                        }
                    }
                }
            }

            if (dialogCaption) {
                dialogCaption.textContent = activeCaption;
            }
            if (dialogPosition) {
                dialogPosition.textContent = String(slideIndex + 1);
            }
        };

        const selectItem = (requestedIndex) => {
            const nextIndex = wrappedIndex(requestedIndex, slides.length);
            currentIndex = nextIndex;
            const activeSlide = slides[currentIndex];

            slides.forEach((slide, index) => {
                if (index !== currentIndex) {
                    resetVideo(slide);
                }
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
                caption.textContent = activeSlide.dataset.caption || "Project media";
            }
            if (position) {
                position.textContent = String(currentIndex + 1);
            }
            syncDialog(activeSlide);
        };

        const selectDialogItem = (direction) => selectItem(currentIndex + direction);

        const startVideo = (slide) => {
            const videoID = slide.dataset.youtubeId;
            const poster = slide.querySelector("[data-gallery-video-open]");
            const player = slide.querySelector("[data-gallery-video-player]");
            if (!videoID || !poster || !player) {
                return;
            }

            const iframe = createVideoIframe(videoID, slide.dataset.caption);

            poster.hidden = true;
            player.hidden = false;
            player.replaceChildren(iframe);
            iframe.focus();
        };

        const startDialogVideo = () => {
            const videoID = dialogVideo?.dataset.youtubeId;
            if (!videoID || !dialogVideoPoster || !dialogVideoPlayer) {
                return;
            }

            const iframe = createVideoIframe(videoID, slides[currentIndex]?.dataset.caption);
            dialogVideoPoster.hidden = true;
            dialogVideoPlayer.hidden = false;
            dialogVideoPlayer.replaceChildren(iframe);
            iframe.focus();
        };

        const openDialog = (opener) => {
            if (!dialog || typeof dialog.showModal !== "function") {
                return false;
            }

            lastOpener = opener;
            const slide = opener.closest("[data-gallery-slide]");
            if (slide) {
                resetVideo(slide);
                selectItem(slides.indexOf(slide));
            }
            dialog.showModal();
            document.body.classList.add("gallery-dialog-open");
            if (closeButton) {
                closeButton.focus();
            }
            return true;
        };

        gallery.querySelectorAll("[data-gallery-direction]").forEach((button) => {
            button.addEventListener("click", () => {
                selectItem(currentIndex + Number(button.dataset.galleryDirection));
            });
        });

        gallery.querySelectorAll("[data-gallery-dialog-direction]").forEach((button) => {
            button.addEventListener("click", () => {
                selectDialogItem(Number(button.dataset.galleryDialogDirection));
            });
        });

        thumbnails.forEach((thumbnail) => {
            thumbnail.addEventListener("click", () => {
                selectItem(Number(thumbnail.dataset.galleryIndex));
            });
        });

        gallery.querySelectorAll("[data-gallery-image-open]").forEach((opener) => {
            opener.addEventListener("click", (event) => {
                if (openDialog(opener)) {
                    event.preventDefault();
                }
            });
        });

        gallery.querySelectorAll("[data-gallery-video-dialog-open]").forEach((opener) => {
            opener.addEventListener("click", () => openDialog(opener));
        });

        gallery.querySelectorAll("[data-gallery-video-open]").forEach((opener) => {
            opener.addEventListener("click", (event) => {
                const slide = opener.closest("[data-gallery-slide]");
                if (!slide) {
                    return;
                }
                event.preventDefault();
                selectItem(slides.indexOf(slide));
                startVideo(slide);
            });
        });

        if (dialogVideoPoster) {
            dialogVideoPoster.addEventListener("click", (event) => {
                if (!dialogVideo?.dataset.youtubeId) {
                    return;
                }
                event.preventDefault();
                startDialogVideo();
            });
        }

        gallery.querySelectorAll("[data-gallery-youtube-thumbnail]").forEach((image) => {
            loadYouTubeThumbnail(image, image.dataset.galleryYoutubeId);
        });

        if (dialog && closeButton) {
            closeButton.addEventListener("click", () => dialog.close());
            dialog.addEventListener("close", () => {
                resetDialogVideo();
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
                if (dialog?.open) {
                    selectDialogItem(-1);
                } else {
                    selectItem(currentIndex - 1);
                }
            } else if (event.key === "ArrowRight") {
                event.preventDefault();
                if (dialog?.open) {
                    selectDialogItem(1);
                } else {
                    selectItem(currentIndex + 1);
                }
            }
        });

        const enableSwipe = (surface, onSwipe) => {
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
                onSwipe(distanceX < 0 ? 1 : -1);
            });
            surface.addEventListener("pointercancel", () => {
                startX = null;
                startY = null;
            });
        };

        enableSwipe(gallery.querySelector("[data-gallery-swipe]"), (direction) => {
            selectItem(currentIndex + direction);
        });
        enableSwipe(gallery.querySelector("[data-gallery-dialog-swipe]"), selectDialogItem);
        selectItem(0);
        gallery.classList.add("project-gallery-ready");
    });
})();
