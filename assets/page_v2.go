package assets

const Page_v2 = `
<!DOCTYPE html>
<html>
<head>
    <title>Site en Go avec JavaScript</title>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/color-thief/2.3.2/color-thief.min.js"></script>
    <style>
        body {
            font-family: Arial, sans-serif;
            height: 100vh;
            margin: 0;
            display: flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;
            background-color: #f0f0f0;
        }

        button {
            margin-top: 20px;
            padding: 10px 15px;
            font-size: 16px;
        }

        img {
            max-width: 500px;
            display: block;
        }
    </style>
</head>
<body onload="loadRandomCat()">
    <img id="randomCat">
    <button id="showRandomCat">Afficher un autre chat</button>

    <script>
        const colorThief = new ColorThief();
        const catImage = document.getElementById('randomCat');

        function loadRandomCat() {
            fetch('/randomcat')
                .then(response => response.json())
                .then(data => {
                    catImage.src = "/catimage?url=" + encodeURIComponent(data[0].url);
                    catImage.onload = function() {
                        const dominantColor = colorThief.getColor(catImage);
                        document.body.style.backgroundColor = 'rgb(' + dominantColor[0] + ',' + dominantColor[1] + ',' + dominantColor[2] + ')';
                    };
                });
        }

        document.getElementById('showRandomCat').addEventListener('click', loadRandomCat);
    </script>
</body>
</html>
`
