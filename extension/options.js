function save_options() {
  const notecollectorUrlInput = document.querySelector('[name="notecollectorUrl"]');

  chrome.storage.sync.set({
    notecollectorUrl: notecollectorUrlInput.value,
  }, () => {
    const saveResult = document.querySelector('.save-result');
    saveResult.textContent = "Saved!";
    setTimeout(() => {
      saveResult.textContent = "";
    }, 1500);
  })

  console.log(e);
}

function restore_options() {
  const notecollectorUrlInput = document.querySelector('[name="notecollectorUrl"]');
  chrome.storage.sync.get({
    notecollectorUrl: 'http://localhost:9872'
  }, (items) => {
    notecollectorUrlInput.value = items.notecollectorUrl;
    notecollectorUrlInput.removeAttribute('disabled');
    notecollectorUrlInput.setAttribute('style', "background-color: white;");
  })
}

document.addEventListener('DOMContentLoaded', restore_options);
document.getElementById('save').addEventListener('click', save_options);
