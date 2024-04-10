function copyToClipboard(button, text) {
    navigator.clipboard.writeText(text)
    .then(() => {
        console.log('Text copied to clipboard:', text);
        const tooltip = new bootstrap.Tooltip(button);
        tooltip.show();
        setTimeout(function() {
            tooltip.hide();
        }, 1000); 
    })
    .catch((err) => {
        console.error('Failed to copy text to clipboard:', err);
    });
}