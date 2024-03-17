sendGetRequest();

// Функция для отправки GET-запроса и обработки ответа
function sendGetRequest() {
    fetch('/ascii-art') // Отправляем GET-запрос на URL '/ascii-art'
        .then(response => response.text()) // Получаем ответ и преобразуем его в текст
        .then(data => { // Обрабатываем полученные данные
            document.getElementById('output').innerText = data; // Выводим данные на страницу
        })
        .catch(error => console.error('Error fetching data:', error)); // В случае ошибки выводим сообщение в консоль
}
function download() {
    fetch('/ascii-art')
        .then(response => response.text()) // Получаем ответ и преобразуем его в текст
        .then(data => {
            // Создаем элемент <a> для загрузки данных как файл
            const a = document.createElement('a');
            // Устанавливаем данные в качестве содержимого ссылки
            a.href = 'data:text/plain;charset=utf-8,' + encodeURIComponent(data);
            // Устанавливаем имя файла для загрузки
            a.download = 'ascii-art.txt';
            // Добавляем элемент <a> в DOM
            document.body.appendChild(a);
            // Нажимаем на ссылку для начала загрузки файла     
            a.click();
            // Удаляем элемент <a> из DOM
            document.body.removeChild(a);
        })
        .catch(error => console.error('Error:', error));
}
