/* 1 定义模块中的AuthController ,并且定义两个action方法 */


Alpaca.MainModule.AuthController = {

    //loginView,  登录页面
    loginViewAction: function () {

        var redirect = Alpaca.Router.getParams(0);

        //视图默认渲染到#content位置，可以通过to对象改变渲染位置
        var view = new Alpaca.View();
        var footer = Alpaca.MainModule.pageFooter();
        view.addChild(footer);
        view.setCaptureTo('body');
        view.ready(function () {
            $('body').addClass('login-container');

            //格式化页面
            LayoutInit();

            //登录按钮 - 调用后台登录接口
            $('.btn-sub-login').unbind('click').click(function () {
                var request = {};
                request.UserName = $('.login-form input[name="login_name"]').val();
                request.PassWd = hex_md5($('.login-form input[name="password"]').val());
                AlpacaAjax({
                    url: g_url + API['admin_auth_login'],
                    data: request,
                    success: function (data) {
                        Notific(data.code);
                        if (data.code == 2009) {
                            //flushCaptcha();
                        }
                        if (data.code == "success") {
                            console.log(redirect);
                            if (redirect) {
                                window.location.replace(decodeURIComponent(redirect));
                                return;
                            }
                            Alpaca.to("#/main/index/index/");
                        }
                    },
                });
            });

            var check_we_times = 0;
            var cur_way = 2;
            $('.btn-login-type').unbind('click').click(function () {
                var way = $(this).data('way');
                cur_way = way;
                $(".login-way").hide();
                if (cur_way == 1) {
                    $(".login-way-email").show();
                }
                if (cur_way == 2) {
                    $(".login-way-wechat").show();
                    check_we_times = 0;
                    checkLogin();
                }
            });

            var checkLogin = function () {
                $.ajax({
                    url: g_url + API['admin_auth_check_wx'],
                    type: "post",
                    success: function (data) {
                        var redirect = Alpaca.Router.getParams(0);
                        if (data.code == 200) {
                            console.log(redirect);
                            Notific(data.code);
                            if (redirect) {
                                window.location.replace(decodeURIComponent(redirect));
                                return;
                            }
                            Alpaca.to("#/main/index/index/");
                        } else {

                            if (cur_way != 2) {
                                return;
                            }

                            check_we_times++;
                            if (check_we_times > 300) {
                                alert("二维码超时，请刷新页面。");
                                return;
                            }
                            setTimeout(function () {
                                checkLogin();
                            }, 1000);
                        }
                    },
                });
            };
            checkLogin();
        });
        return view;
    },

    //注销-调用服务器接口
    logoutAction: function () {
        AlpacaAjax({
            url: g_url + API['admin_auth_logout'],
            newSuccess: function (data) {
                if (data.code == "success") {
                    AlpacaCache.clear();
                    window.location.href = "/web/admin/#/main/auth/loginView";
                } else {
                    AlpacaCache.clear();
                    alert(data.code);
                }
            },
        });
    },

    //获取用户个人信息 - 页面
    myInfoViewAction: function () {
        var userInfo = Alpaca.MainModule.getUserInfo(true);
        var view = new Alpaca.MainModule.pageView({data: userInfo['member']});
        return view;
    },

    //重置密码 - 页面
    resetPasswordViewAction: function () {
        var view = new Alpaca.MainModule.pageView();
        return view;
    },

    //RegView,  登录页面
    regViewAction: function () {

        //视图默认渲染到#content位置，可以通过to对象改变渲染位置
        var view = new Alpaca.View();
        var footer = Alpaca.MainModule.pageFooter();
        view.addChild(footer);
        view.setCaptureTo('body');
        view.ready(function () {
            $('body').addClass('login-container');

            //格式化页面
            LayoutInit();
        });
        return view;
    },

    //RegView,  登录页面
    forgetViewAction: function () {

        //视图默认渲染到#content位置，可以通过to对象改变渲染位置
        var view = new Alpaca.View();
        var footer = Alpaca.MainModule.pageFooter();
        view.addChild(footer);
        view.setCaptureTo('body');
        view.ready(function () {
            $('body').addClass('login-container');

            //格式化页面
            LayoutInit();
        });
        return view;
    },

};