function getCurrentTab(callback) {
  let queryOptions = { active: true, lastFocusedWindow: true };
  chrome.tabs.query(queryOptions, ([tab]) => {
    if (chrome.runtime.lastError)
    console.error(chrome.runtime.lastError);
    // `tab` will either be a `tabs.Tab` instance or `undefined`.
    callback(tab);
  });
}

function main() {
  getCurrentTab((tab) => {
    const linkInput = document.querySelector('input[name="Link"]');
    linkInput.setAttribute('value', tab.url);
  });
}

function save_note() {
  const linkInput = document.querySelector('[name="Link"]');
  const textInput = document.querySelector('[name="Text"]');
  const saveResult = document.querySelector('.save-result');

  chrome.storage.sync.get({
    notecollectorUrl: 'http://localhost:9872'
  }, (items) => {
    fetch(`${items.notecollectorUrl}/notes`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        Link: linkInput.value,
        Text: textInput.value
      })
    }).then(() => {
      saveResult.setAttribute('style', 'color: gray;');
      saveResult.textContent = "Saved!";
      setTimeout(() => {
        saveResult.textContent = "";
        window.close();
      }, 750);
    })
    .catch((err) => {
      saveResult.textContent = `Error! ${err}`;
      saveResult.setAttribute('style', 'color: red;');
    });
  });
}

document.addEventListener('DOMContentLoaded', main);
document.getElementById('save').addEventListener('click', save_note);
