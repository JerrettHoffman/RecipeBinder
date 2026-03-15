const newlineRegex = /\r?\n/;
const validBulletRegex = /([\p{L}\p{N}].*)/ui;
const sectionRegex = /^## /;

const formattedIngredients = document.getElementById("formatted-ingredients");
const ingredientArea = document.getElementById("ingredients");
ingredientArea.addEventListener("input", (e) => {
	formattedIngredients.innerHTML = ingredientArea.value.split(newlineRegex).map((line) => {
		// Leave lines that start with '## ' alone
		if (sectionRegex.test(line)) {
			return line;
		}

		// Match the first alphanumeric character and everything after
		const bulletText = validBulletRegex.exec(line);
		if (bulletText === null) {
			// Lines without any alphanumeric characters are skipped
			return "";
		} else {
			return "* " + bulletText[0];
		}
	}).filter(Boolean).join("\n");
});
