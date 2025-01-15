let output = document.getElementById("output");
let tab = document.getElementById("output_tab");
let input = document.getElementById("input")
let run_button = document.getElementById("run")
output.value = "";

document.addEventListener("keydown", (e) => {
	if (e.key == "Enter" && e.ctrlKey) {
		run_button.click()
	}
}) 
