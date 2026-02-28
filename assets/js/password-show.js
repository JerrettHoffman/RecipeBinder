const passwordShowBox = document.getElementById("password-show");
const passwordInput = document.getElementById("password");

passwordShowBox.addEventListener("change", () => {
	if (passwordInput.type === "text") {
		passwordInput.type = "password";
	} else {
		passwordInput.type = "text";
	}
})
