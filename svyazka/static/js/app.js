document.getElementById('fetchBtn').addEventListener('click', async () => {
    try {
        const response = await fetch('/api/data');
        const data = await response.json();
        document.getElementById('output').innerHTML = 
            `<p>Ответ: ${data.message}</p>`;
    } catch (error) {
        console.error('Error:', error);
    }
});