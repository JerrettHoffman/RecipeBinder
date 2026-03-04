const addIngredientButton = document.getElementById("add-ingredient");
const ingredientsContainer = document.getElementById("ingredients-container");
const ingredientCount = document.getElementById("ingredient-count");

function assignIngredientNumber(ingredientFormGroup, index) {
	// Rename the inputs
	let ingredientInput = ingredientFormGroup.querySelector("input");
	ingredientInput.id = "ingredient-" + index;
	ingredientInput.name = "ingredient-" + index;
	// Rename the buttons
	let ingredientRemoveButton = ingredientFormGroup.querySelector("button");
	ingredientRemoveButton.id = "remove-ingredient-" + index;
}

function shouldViewTransition() {
	return document.startViewTransition && window.matchMedia('(prefers-reduced-motion: no-preference)').matches;
}

function setupRemoval(ingredientFormGroup) {
	let button = ingredientFormGroup.querySelector("button");
	button.addEventListener("click", () => {
		function assignNumbers() {
			let ingredientList = document.querySelectorAll(".ingredient-group");
			ingredientList.forEach((ingredientFormGroup, i) => assignIngredientNumber(ingredientFormGroup, i));
			ingredientCount.value = ingredientList.length;
		}

		if (!shouldViewTransition()) {
			ingredientFormGroup.remove();
			assignNumbers();
		} else {
			const transition = document.startViewTransition(() => ingredientFormGroup.remove());
			transition.finished.then(assignNumbers);
		}
	});
}

addIngredientButton.addEventListener("click", () => {
	// Test to see if the browser supports the HTML template element by checking
	// for the presence of the template element's content attribute.
	if ("content" in document.createElement("template")) {
		const template = document.querySelector("#ingredient-template");
		const clone = document.importNode(template.content, true);

		let ingredientGroup = clone.children[0];
		setupRemoval(ingredientGroup);

		let ingredientList = document.querySelectorAll(".ingredient-group");
		assignIngredientNumber(ingredientGroup, ingredientList.length);
		ingredientCount.value = ingredientList.length + 1;

		// Add the clone to the container
		function insertClone() {
			ingredientsContainer.insertBefore(clone, addIngredientButton);
		}

		if (!shouldViewTransition()) {
			insertClone();
		} else {
			const transition = document.startViewTransition(() => insertClone());
		}
	}
});

// Set up removal for any ingredients added by the golang template
let ingredientsList = document.querySelectorAll(".ingredient-group");
ingredientsList.forEach((ingredientGroup, i) => {
	setupRemoval(ingredientGroup)
	assignIngredientNumber(ingredientGroup, i);
});
