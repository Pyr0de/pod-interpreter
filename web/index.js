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

input.addEventListener('keydown', function(e) {
  if (e.key == 'Tab') {
    e.preventDefault();
    var start = this.selectionStart;
    var end = this.selectionEnd;

    // set textarea value to: text before caret + tab + text after caret
    this.value = this.value.substring(0, start) +
      "\t" + this.value.substring(end);

    // put caret at right position again
    this.selectionStart =
      this.selectionEnd = start + 1;
  }
});
