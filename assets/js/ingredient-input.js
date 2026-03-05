const ingredientCount = document.getElementById("ingredient-count");

function assignIngredientNumber(ingredientFormGroup, index) {
	// Rename the inputs
	let ingredientInput = ingredientFormGroup.querySelector("input");
	ingredientInput.id = "ingredient-" + index;
	ingredientInput.name = "ingredient-" + index;
}

function shouldViewTransition() {
	return document.startViewTransition && window.matchMedia('(prefers-reduced-motion: no-preference)').matches;
}

function removeIngredient(ingredientsContainer, removeButton) {
	let groupToRemove = removeButton.parentNode;
	function assignNumbers() {
		let ingredientList = ingredientsContainer.querySelectorAll(".ingredient-group");
		ingredientList.forEach((ingredientFormGroup, i) => assignIngredientNumber(ingredientFormGroup, i));
		ingredientCount.value = ingredientList.length;
	}

	if (!shouldViewTransition()) {
		groupToRemove.remove();
		assignNumbers();
	} else {
		const transition = document.startViewTransition(() => groupToRemove.remove());
		transition.finished.then(assignNumbers);
	}
}

function addIngredient(ingredientsContainer, addIngredientButton) {
	// Test to see if the browser supports the HTML template element by checking
	// for the presence of the template element's content attribute.
	if ("content" in document.createElement("template")) {
		const template = document.querySelector("#ingredient-template");
		const clone = document.importNode(template.content, true);

		let ingredientGroup = clone.children[0];

		let ingredientList = ingredientsContainer.querySelectorAll(".ingredient-group");
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
}

let ingredientContainers = document.querySelectorAll(".ingredient-container");
ingredientContainers.forEach((container, i) => {
	container.addEventListener("command", e => {
		if (e.command === "--remove-ingredient") {
			removeIngredient(e.target, e.source);
		} else if (e.command === "--add-ingredient") {
			addIngredient(e.target, e.source);
		}
	});
});
