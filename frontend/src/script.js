const comments = document.querySelectorAll(".post-comment");
const commentElements = document.querySelectorAll(".CommentPost");
const BLOCK_DISPLAY = "block";
const NONE_DISPLAY = "none";
const passwordToggles = document.querySelectorAll(".toggle");

const isElementHidden = function(element) {
  const { display } = element.style;
  return display === "" || display === NONE_DISPLAY;
};

const toggleCommentDisplay = () => {
  commentElements.forEach((element) => {
    element.style.display = isElementHidden(element) ? BLOCK_DISPLAY : NONE_DISPLAY;
  });
};

document.addEventListener('DOMContentLoaded', () => {

  // COMMENT DISPLAY TOGGLE
  comments.forEach((comment) => {
    comment.addEventListener("click", toggleCommentDisplay);
  });

  // PASSWORD EYE
  passwordToggles.forEach((toggle) => {
    toggle.addEventListener('click', () => {
      const passwordInput = toggle.previousElementSibling;
      passwordInput.type = passwordInput.type === 'password' ? 'text' : 'password';
    });
  });

  // TEXTAREA HEIGHT AUTO-EXPAND
  const textareaElements = document.querySelectorAll('textarea');
  textareaElements.forEach((textarea) => {
    textarea.addEventListener('input', () => {
      textarea.style.height = 'auto';
      textarea.style.height = `${textarea.scrollHeight}px`;
    });
  });
});
