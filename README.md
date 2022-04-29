# Final Project 1 (TODO's APP)

Project ini adalah backend service yang interfacenya berupa API Endpoint, data yang di sediakan ke API
akan dimasukkan ke slice of struct yang dimana akan membuat data tersebut te-reset setiap runtime.

## API Endpoint
- {GET} /todos/ (mengambil semua data yang di keluarkan dalam bentuk array of json)
- {GET} /todos/:id (mengambil data spesifik berdasarkan id yang di keluarkan dalam bentuk json)
- {POST} /todos/ (memberikan / menulis data ke dalam backend yang merupakan todos baru dengan id yang auto-increment)
- {PUT} /todos/:id (mengubah konten json dri sebuah todo berdasarkan id)
- {DELETE} /todos/:id (menghapus konten json / todos berdasarkan id)

## Anggota Kelompok 
- Muhammad Daffa Haryadi 
- I Gede Diva
- Denada Ramschie

## Resources
- [How to Design better apis](https://r.bluethl.net/how-to-design-better-apis) 
- [swaggo docs](https://github.com/swaggo/swag) 
