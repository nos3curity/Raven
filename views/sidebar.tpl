<!doctype html>
<html lang="en">
<head>
    <title>Raven - Dashboard</title>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="/static/bootstrap/css/bootstrap.min.css">
    <script src="/static/bootstrap/js/bootstrap.min.js"></script>
    <link rel="stylesheet" href="/static/css/custom.css">
</head>
<body>
<div class="fluid-container">
    <div class="row g-0">
        <div class="col-12 col-sm-6 col-md-4 col-lg-3 col-xl-3 p-0 cs-bg-gray">
            <div class="d-flex flex-column align-items-center align-items-sm-start min-vh-100 p-0 cs-bg-gray pe-3 ps-3">
                <!-- <nav class="navbar p-0"> -->
                    <div class="container-fluid p-3">
                        <!-- <ul class="navbar-nav cs-bg-gray col-12 col-md-4 col-lg-3 col-xl-3 ps-4 pt-4 pb-4 pe-4"> -->
                            <h2 class="fw-bold">
                                Raven
                            </h2>
                            <div>
                                <a href="/">
                                    <button type="button" class="btn w-100 text-start rounded-0 cs-button-active cs-text-white mb-2">Home</button>
                                </a>
                                <a href="/teams">
                                    <button type="button" class="btn w-100 text-start rounded-0 cs-text-black mb-2">Team Setup</button>
                                </a>
                                <a href="/uploads">
                                    <button type="button" class="btn w-100 text-start rounded-0 cs-text-black mb-2">Uploads</button>
                                </a>
                                <a href="/profile">
                                    <button type="button" class="btn w-100 text-start rounded-0 cs-text-black mb-2">Profile</button>
                                </a>
                            </div>
                            <br>
                            <br>
                            <h4 class="fw-bold">
                                Teams
                            </h4>
                            <div>
                            {{range .teams}}
                                <a href="/teams/{{.Id}}">
                                    <button type="button" class="btn w-100 text-start rounded-0 cs-text-black mb-2">{{.Name}}</button>
                                </a>
                            {{end}}
                            </div>
                        <!-- </ul> -->
                    </div>
                <!-- </nav>     -->
            </div>
        </div>
        {{ .LayoutContent }}
    </div>
</div>


</div>





</body>
</html>
