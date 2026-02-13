const addIngredientButton = document.getElementById("add-ingredient");
const ingredientsContainer = document.getElementById("ingredients-container");
const ingredientCount = document.getElementById("ingredient-count");

function reassignIngredientIds(ingredientList) {
	ingredientList.forEach((ingredientFormGroup, i) => {
		// Rename the inputs
		let ingredientInput = ingredientFormGroup.querySelector("input");
		ingredientInput.id = "ingredient-" + i;
		ingredientInput.name = "ingredient-" + i;
		// Rename the buttons
		let ingredientRemoveButton = ingredientFormGroup.querySelector("button");
		ingredientRemoveButton.id = "remove-ingredient-" + i;
	});
}

function setupRemoval(ingredientFormGroup) {
	let button = ingredientFormGroup.querySelector("button");
	button.addEventListener("click", () => {
		ingredientFormGroup.remove();
		let ingredientList = document.querySelectorAll(".ingredient-group");
		reassignIngredientIds(ingredientList);
		ingredientCount.value = ingredientList.length;
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

		// Add the clone to the container
		ingredientsContainer.insertBefore(clone, addIngredientButton);

		let ingredientList = document.querySelectorAll(".ingredient-group");
		reassignIngredientIds(ingredientList);
		ingredientCount.value = ingredientList.length;
	}
});

// Set up removal for any ingredients added by the golang template
let ingredientsList = document.querySelectorAll(".ingredient-group");
ingredientsList.forEach((ingredientGroup) => {
	setupRemoval(ingredientGroup)
});
