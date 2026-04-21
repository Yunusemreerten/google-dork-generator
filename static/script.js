function copyQuery(text) {
    navigator.clipboard.writeText(text)
        .then(() => {
            alert("Query copied");
        })
        .catch(() => {
            alert("Copy failed");
        });
}

function openGoogle(query) {
    const url = "https://www.google.com/search?q=" + encodeURIComponent(query);
    window.open(url, "_blank");
}

document.addEventListener("DOMContentLoaded", function () {
    const copyButtons = document.querySelectorAll(".copy-btn");
    const googleButtons = document.querySelectorAll(".google-btn");

    copyButtons.forEach(button => {
        button.addEventListener("click", function () {
            const query = this.getAttribute("data-query");
            copyQuery(query);
        });
    });

    googleButtons.forEach(button => {
        button.addEventListener("click", function () {
            const query = this.getAttribute("data-query");
            openGoogle(query);
        });
    });
});