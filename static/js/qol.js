// Initializing Bootstrap's tooltips
// Used for "Copied!" pop-up
var tooltipTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="tooltip"]'))
var tooltipList = tooltipTriggerList.map(function (tooltipTriggerEl) {
  return new bootstrap.Tooltip(tooltipTriggerEl)
})



// Get all dropdown items
var dropdownItems = document.querySelectorAll('.dropdown-item');

// Loop through each dropdown item and add click event listener
dropdownItems.forEach(function(item) {
    item.addEventListener('click', function() {
        // Get the image URL of the clicked dropdown item
        var imageURL = item.querySelector('img').getAttribute('src');
           
        // Update the image URL of the span with id 'dropdown-title'
        document.getElementById('dropdown-title').innerHTML = '<img src="' + imageURL + '" style="width: 25px;">';
    });
});


// Get all elements whose IDs start with "accordion-button-"
const accordionButtons = document.querySelectorAll('[id^="accordion-button-"]');

// Iterate thru all accordion buttons
accordionButtons.forEach(function(accordionButton) {

    // Extract accordion button index
    const index = accordionButton.id.split('-').pop();
    console.log(index);

    // Construct accordionHeadingId based on hardcoded "accordion-heading-" prefix and index
    const accordionHeadingId = 'accordion-heading-' + index;
    const accordionItemId = 'accordion-item-' + index;

    console.log(accordionHeadingId);
    console.log(accordionItemId);

    // Get accordion heading and accordion item objects
    const accordionHeading = document.getElementById(accordionHeadingId);
    const accordionItem = document.getElementById(accordionItemId);

    // On mouseover event, make style changes
    accordionButton.addEventListener('mouseover', function() {
        if (accordionHeading && accordionItem) {
            accordionHeading.classList.remove('cs-bg-gray'); 
            accordionHeading.style.backgroundColor = '#adadad';

            accordionItem.classList.remove('cs-bg-gray');
            accordionItem.style.backgroundColor = '#adadad';
        }
    });

    // On mouseout event, make style changes
    accordionButton.addEventListener('mouseout', function() {
        if (accordionHeading && accordionItem) {
            accordionHeading.classList.add('cs-bg-gray'); 
            accordionHeading.style.backgroundColor = ''; 

            accordionItem.classList.add('cs-bg-gray');
            accordionItem.style.backgroundColor = '';
        }
    });
});


// ==== copyToClipboard =======================================================
//
// This function copies the input "text" to the user's clipboard. It also 
// enables the calling button's tooltip temporarily. Usually, the tooltip will
// contain text such as "Copied!" to verify to the user that the text was
// successfully copied. 
//
// Input:
//      copyButton   -- a reference to the calling button object
//      text [IN]    -- a string value to be copied
//
// ============================================================================
function copyToClipboard(copyButton, text) {
    // Copy "text" to system clipboard
    navigator.clipboard.writeText(text)
    .then(() => {
        console.log('Text copied to clipboard:', text);

        // Get tooltip object based on copyButton
        const tooltip = new bootstrap.Tooltip(copyButton);
        // Show tooltip ("Copied!" -- or similar verified text) 
        tooltip.show();
        setTimeout(function() {
            tooltip.hide();
        }, 1000); 
    })
    .catch((err) => {
        console.error('Failed to copy text to clipboard:', err);
    });
}



