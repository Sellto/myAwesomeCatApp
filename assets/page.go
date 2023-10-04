package assets

const Page = `
<!DOCTYPE html>
<html>
<head>
    <title>Site en Go avec JavaScript</title>
</head>
<body>
    <button id="changeBackground">Changer le fond d'écran</button>
    <button id="showRandomCat">Afficher un chat aléatoire</button>
    <img id="randomCat" style="display: none; max-width: 500px; margin-top: 20px;">

    <script>
        document.getElementById('changeBackground').addEventListener('click', function() {
            document.body.style.backgroundColor = '#' + Math.floor(Math.random()*16777215).toString(16);
        });

        document.getElementById('showRandomCat').addEventListener('click', function() {
            fetch('https://api.thecatapi.com/v1/images/search')
                .then(response => response.json())
                .then(data => {
                    const catImage = document.getElementById('randomCat');
                    catImage.src = data[0].url;
                    catImage.style.display = 'block';
                });
        });
    </script>
</body>
</html>
`
