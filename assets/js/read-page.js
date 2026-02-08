const dialog = document.getElementById("ingredient-popup");
const openButton = document.getElementById("open-ingredient-popup");
const closeButton = document.getElementById("close-ingredient-popup");

openButton.addEventListener("click", () => {
	dialog.showModal();
});

closeButton.addEventListener("click", () => {
	dialog.close();
});
