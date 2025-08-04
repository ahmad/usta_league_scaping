chrome.runtime.onMessage.addListener((request, sender, sendResponse) => {
  if (request.action === "postData") {
    console.log(request.payload);

    fetch('http://localhost:8080/', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(request.payload),
    })
    .then(response => response.json())
    .then(data => {
      console.log('Success:', data);
      sendResponse({status: "success", data: data});
    })
    .catch((error) => {
      console.error('Error:', error);
      sendResponse({status: "error", error: error});
    });
    return true; // Indicates that the response is sent asynchronously
  }
});