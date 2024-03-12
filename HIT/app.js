const axios = require("axios");
const readline = require('readline');

const token = '1984261227:AAH3EN2UBH1ukOMDpzp2udWAXTrYsc4sfG0';
const chatId = "1450191429";

// Membuat interface readline
const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

// Fungsi untuk mengirim pesan ke bot Telegram
function sendMessage(message) {
    axios.post(`https://api.telegram.org/bot${token}/sendMessage`, {
        chat_id: chatId,
        text: message
    })
    .catch((error) => {
        console.error('Error:', error);
    });
}

// Fungsi untuk meminta masukan dari pengguna secara berulang
function askForInput() {
    rl.question('Masukkan pesan: ', (input) => {
        // Jika input pengguna adalah ##
        if (input === '##') {
            rl.close(); // Menutup interface readline
        } else {
            // Mengirim pesan ke bot Telegram
            sendMessage(input);
            // Meminta masukan lagi
            askForInput();
        }
    });
}

// Memulai proses meminta masukan dari pengguna
askForInput();
