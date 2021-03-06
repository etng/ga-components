package theme1

var List = map[string]string{"login/theme1": `{{define "login_theme1"}}
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="utf-8">
        <link rel="apple-touch-icon" sizes="76x76" href="{{.UrlPrefix}}/assets/img/apple-icon.png">
        <link rel="icon" type="image/png" href="{{.UrlPrefix}}/assets/img/favicon.png">
        <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
        <title>Login</title>
        <meta content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0, shrink-to-fit=no"
              name="viewport">
        <link href="https://fonts.googleapis.com/css?family=Montserrat:400,700,200" rel="stylesheet">
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/latest/css/font-awesome.min.css">
        <link href="{{.UrlPrefix}}/assets/css/bootstrap.min.css" rel="stylesheet">
        <link href="{{.UrlPrefix}}/assets/css/now-ui-kit.css?v=1.1.0" rel="stylesheet">
        <link href="{{.UrlPrefix}}/assets/css/demo.css" rel="stylesheet">
        <link rel="canonical" href="">
        <meta name="keywords" content="">
        <meta name="description" content="">
    </head>

    <body class="login-page sidebar-collapse">
    <div class="page-header" filter-color="orange">
        <div class="page-header-image" style="background-image:url({{.UrlPrefix}}/assets/img/login.jpg)"></div>
        <div class="container">
            <div class="col-md-4 content-center">
                <div class="card card-login card-plain">
                    <form action="{{.UrlPrefix}}/signin" method="post" id="sign-up-form">
                        <div class="header header-primary text-center">
                            <div class="logo-container">
                                <img src="assets/img/now-logo.png" alt="">
                            </div>
                        </div>
                        <div class="content">
                            <div class="input-group form-group-no-border input-lg">
                                <span class="input-group-addon">
                                    <i class="now-ui-icons users_circle-08"></i>
                                </span>
                                <input type="text" id="username" class="form-control" placeholder="username" autocomplete="off">
                            </div>
                            <div class="input-group form-group-no-border input-lg">
                                <span class="input-group-addon">
                                    <i class="now-ui-icons text_caps-small"></i>
                                </span>
                                <input type="password" id="password" placeholder="password" class="form-control" autocomplete="off">
                            </div>
                        </div>
                        <div class="footer text-center">
                            <button class="btn btn-primary btn-round btn-lg btn-block">Login</button>
                        </div>
                    </form>
                </div>
            </div>
        </div>
        <footer class="footer">
            <div class="container">
                <div class="copyright">
                    © 2019, created by <a href="http://www.go-admin.cn/">GoAdmin.</a>
                </div>
            </div>
        </footer>
    </div>

    <script src="{{.UrlPrefix}}/assets/js/core/jquery.3.2.1.min.js" type="text/javascript"></script>
    <script src="{{.UrlPrefix}}/assets/js/core/popper.min.js" type="text/javascript"></script>
    <script src="{{.UrlPrefix}}/assets/js/core/bootstrap.min.js" type="text/javascript"></script>
    <script src="{{.UrlPrefix}}/assets/js/plugins/bootstrap-switch.js"></script>
    <script src="{{.UrlPrefix}}/assets/js/plugins/nouislider.min.js" type="text/javascript"></script>
    <script src="{{.UrlPrefix}}/assets/js/plugins/bootstrap-datepicker.js" type="text/javascript"></script>
    <script src="{{.UrlPrefix}}/assets/js/plugins/jquery.sharrre.js" type="text/javascript"></script>
    <script src="{{.UrlPrefix}}/assets/js/now-ui-kit.js?v=1.1.0" type="text/javascript"></script>
    <script>
        $("#sign-up-form").submit(function (e) {
            e.preventDefault();
            $.ajax({
                dataType: 'json',
                type: 'POST',
                url: '{{.UrlPrefix}}/signin',
                async: 'true',
                data: {
                    'username': $("#username").val(),
                    'password': $("#password").val()
                },
                success: function (data) {
                    location.href = data.data.url
                },
                error: function (data) {
                    alert("登录失败");
                }
            });
        });
    </script>

    </body>
    </html>
{{end}}`}
