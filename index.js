document.getElementById('enviar').addEventListener('click', function() {
  
    const url = 'http://localhost:8080/api/consulta';
    const data = {
      producto: document.getElementById('prd').value
    };
    console.log(data);
    console.log(JSON.stringify(data));
    fetch(url, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
    .then(response => response.json())
    .then(result => console.log(result))
  });