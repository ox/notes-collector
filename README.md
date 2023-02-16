# Notes Collector

Notes Collector is a server that receives an object with a URL and some text from a chrome extension and saves it in a file. It's useful for collecting links you find as you browse.

## Setting up the server

Download the source and run `./run.sh`.

## Installing the Chrome Extension

Enter Developer Mode in Chrome and load the unpacked extension in the `extension` folder. After installing the extension you can configure it to point to your local note-collector service. The default options should work just fine.

Navigate to some page, click the extension, write some text about the page, and hit "Save." If everything was set up correctly you should see "saved!" and then the Chrome popup will disappear. The notesfile should have a new entry in it with the link that you saved.
