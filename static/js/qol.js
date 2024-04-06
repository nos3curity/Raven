function copyToClipboard(textToCopy) {
    var iconElement = event.target; // Get the clicked element

    navigator.clipboard.writeText(textToCopy).then(function() {
        // Change to check icon
        iconElement.classList.remove('fa-copy');
        iconElement.classList.add('fa-check');

        // Set timeout to revert back to the copy icon
        setTimeout(function() {
            iconElement.classList.remove('fa-check');
            iconElement.classList.add('fa-copy');
        }, 1000); // 1 second delay
    }, function(err) {
        console.error('Error in copying text: ', err);
    });
}
