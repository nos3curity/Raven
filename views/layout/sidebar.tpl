<!doctype html>
<html lang="en">
<head>
    <title>Raven - Dashboard</title>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="/static/bootstrap/css/bootstrap.min.css">

    <!-- START: FOR TOOLTIPS -->
    <script src="https://cdn.jsdelivr.net/npm/@popperjs/core@2.11.6/dist/umd/popper.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.2.0-alpha1/dist/js/bootstrap.bundle.min.js"></script>
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js"></script>
    <!-- END: FOR TOOLTIPS -->

    <script src="/static/bootstrap/js/bootstrap.min.js"></script>
    <script src="/static/js/qol.js"></script>
    <link rel="stylesheet" href="/static/css/custom.css">
</head>
<body>
<script src="http://localhost:8000/copilot/index.js"></script>
<script>
    window.mountChainlitWidget({
    chainlitServer: "http://localhost:8000",
    button: {
        style: {
          bgcolor: "#929190", // Set the background color of the button
          bgcolorHover: "#bdbdbc"
        }
      }
    });
</script>
<div class="fluid-container">
    <div class="row g-0">
        <div class="col-12 col-sm-5 col-md-4 col-lg-3 col-xl-3 p-0 cs-bg-gray">
            <div class="d-flex flex-column align-items-center align-items-sm-start min-vh-100 p-0 cs-bg-gray pe-3 ps-3">
                <!-- <nav class="navbar p-0"> -->
                    <div class="container-fluid p-3">
                        <!-- <ul class="navbar-nav cs-bg-gray col-12 col-md-4 col-lg-3 col-xl-3 ps-4 pt-4 pb-4 pe-4"> -->
                            <h2 class="fw-bold">
                                <a class="cs-text-black text-decoration-none" href="/">Raven</h2></a>
                            <div>
                                <!-- <a href="/" onclick="setActiveButton('homeBtn')">
                                    <button id="homeBtn" type="button" class="cs-button btn w-100 text-start rounded-0 mb-2">Home</button>
                                </a> -->
                                <a href="/teams" onclick="setActiveButton('teamsBtn')">
                                    <button id="teamsBtn" type="button" class="cs-button btn w-100 text-start rounded-0 mb-2 py-2">Team Setup</button>
                                </a>
                                <a href="/uploads" onclick="setActiveButton('uploadsBtn')">
                                    <button id="uploadsBtn" type="button" class="cs-button btn w-100 text-start rounded-0 mb-2 py-2">Uploads</button>
                                </a>
                                <a href="/profile" onclick="setActiveButton('profileBtn')">
                                    <button id="profileBtn" type="button" class="btn w-100 text-start rounded-0 cs-text-black mb-2 py-2">Profile</button>
                                </a>
                                <a href="/loot" onclick="setActiveButton('lootBtn')">
                                    <button id="lootBtn" type="button" class="btn w-100 text-start rounded-0 cs-text-black mb-2 py-2">Loot</button>
                                </a>
                            </div>
                            <br>
                            <br>
                            <h4 class="fw-bold">
                                Teams
                            </h4>
                            <div>
                                {{ range .teams }}
                                <a href="/teams/{{ .Id }}" onclick="setActiveButton('team{{ .Id }}Btn')">
                                    <button id="team{{ .Id }}Btn" type="button" class="btn w-100 text-start rounded-0 cs-text-black mb-2 py-2">{{ .Name }}</button>
                                </a>
                                {{ end }}
                            </div>
                        <!-- </ul> -->
                    </div>
                <!-- </nav>     -->
            </div>
        </div>
        <div class="col-12 col-sm-7 col-md-8 col-lg-9 col-xl-9 p-0 cs-bg-black">
            {{ .LayoutContent }}
        </div>
    </div>
</div>


</div>

<script>
    // Function to set the active button and store it in localStorage
    function setActiveButton(buttonId) {
            // Remove the 'cs-button-active' class from all buttons
            document.querySelectorAll('button').forEach(function(btn) {
                btn.classList.remove('cs-button-active');
            });
    
            // Add the 'cs-button-active' class to the clicked button
            document.getElementById(buttonId).classList.add('cs-button-active');
    
            // Store the active button ID in localStorage
            localStorage.setItem('activeButton', buttonId);
        }
    
        // Function to set the active button on page load
        function setActiveButtonOnLoad() {
            // Retrieve the active button ID from localStorage
            var activeButtonId = localStorage.getItem('activeButton');
    
            // If an active button ID is found, set the active button
            if (activeButtonId) {
                setActiveButton(activeButtonId);
            }
        }
    
        // Set the active button on page load
        setActiveButtonOnLoad();
    </script>




</body>
</html>
