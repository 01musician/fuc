// Get references to HTML elements
const button = document.getElementById('myButton');
const messageDisplay = document.getElementById('messageDisplay');

// Event listener for the button click event
button.addEventListener('click', function () {
    // Update the message when the button is clicked
    showMessage('Button clicked!');
});

// Function to display messages
function showMessage(message) {
    messageDisplay.textContent = message;
}
