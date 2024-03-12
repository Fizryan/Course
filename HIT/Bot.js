const TelegramBot = require('node-telegram-bot-api');
const fs = require('fs');

// Token bot Telegram
const botToken = '1984261227:AAH3EN2UBH1ukOMDpzp2udWAXTrYsc4sfG0';

// Buat bot Telegram menggunakan token
const bot = new TelegramBot(botToken, { polling: true });

// Mendefinisikan nama file untuk menyimpan pesan
const fileName = 'pesan_pengguna.txt';

// Mendengarkan pesan dari pengguna
bot.on('message', (msg) => {
    const chatId = msg.chat.id;
    const messageText = msg.text;
    const userId = msg.from.id; // Mendapatkan ID pengguna

    // Memproses pesan dari pengguna
    if (messageText) {
        // Lakukan sesuatu berdasarkan isi pesan yang diterima
        if (messageText === '/start') {
            fs.appendFile(fileName, `User ID: ${userId} - ${messageText}\n`, (err) => {
                if (err) {
                    console.error(err);
                }
            });
        } else {
            // Menyimpan pesan teks ke dalam file teks
            fs.appendFile(fileName, `User ID: ${userId} - ${messageText}\n`, (err) => {
                if (err) {
                    console.error(err);
                }
            });
        }
    }
});