// On "submit" event, do the following:
document.getElementById('commentForm').addEventListener('submit', function(event) {
    var formData = new FormData(this);

    fetch('/comments/add', {
        method: 'POST',
        body: formData,
    })
    .then(response => response.json())
    .then(data => {
        // Update the comments list with the new comment
        var commentsList = document.getElementById('commentsList');
        var newCommentElement = document.createElement('li');
        newCommentElement.textContent = data.createdAt + '/' + data.username + '/' + data.text;
        commentsList.appendChild(newCommentElement);

        // Clear the textarea
        document.querySelector('textarea[name="comment"]').value = '';
    })
    .catch(error => {
        console.error('Error:', error);
    });
});