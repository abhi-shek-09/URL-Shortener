// Function to shorten URL (POST request)
async function shortenURL() {
    const originalURL = document.getElementById("urlInput").value;

    const response = await fetch("http://localhost:8080/shorten", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ original_url: originalURL })
    });

    const data = await response.json();
    
    if (response.ok) {
        document.getElementById("shortenedLink").innerHTML = 
            `<a href="${data.short_url}" target="_blank">${data.short_url}</a>`;
    } else {
        alert("Error: " + data.error);
    }
}


