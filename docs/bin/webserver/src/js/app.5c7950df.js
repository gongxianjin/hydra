(function(t){function e(e){for(var r,a,u=e[0],s=e[1],l=e[2],f=0,d=[];f<u.length;f++)a=u[f],Object.prototype.hasOwnProperty.call(o,a)&&o[a]&&d.push(o[a][0]),o[a]=0;for(r in s)Object.prototype.hasOwnProperty.call(s,r)&&(t[r]=s[r]);c&&c(e);while(d.length)d.shift()();return i.push.apply(i,l||[]),n()}function n(){for(var t,e=0;e<i.length;e++){for(var n=i[e],r=!0,u=1;u<n.length;u++){var s=n[u];0!==o[s]&&(r=!1)}r&&(i.splice(e--,1),t=a(a.s=n[0]))}return t}var r={},o={app:0},i=[];function a(e){if(r[e])return r[e].exports;var n=r[e]={i:e,l:!1,exports:{}};return t[e].call(n.exports,n,n.exports,a),n.l=!0,n.exports}a.m=t,a.c=r,a.d=function(t,e,n){a.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:n})},a.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},a.t=function(t,e){if(1&e&&(t=a(t)),8&e)return t;if(4&e&&"object"===typeof t&&t&&t.__esModule)return t;var n=Object.create(null);if(a.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var r in t)a.d(n,r,function(e){return t[e]}.bind(null,r));return n},a.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return a.d(e,"a",e),e},a.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},a.p="/";var u=window["webpackJsonp"]=window["webpackJsonp"]||[],s=u.push.bind(u);u.push=e,u=u.slice();for(var l=0;l<u.length;l++)e(u[l]);var c=s;i.push([0,"chunk-vendors"]),n()})({0:function(t,e,n){t.exports=n("56d7")},"56d7":function(t,e,n){"use strict";n.r(e);n("e260"),n("e6cf"),n("cca6"),n("a79d");var r=n("a026"),o=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{attrs:{id:"vapp"}},[n("Row",[n("i-col",{attrs:{span:"24"}},[n("Menu",{attrs:{mode:"horizontal",theme:"dark","active-key":"1"}},[n("Row",{attrs:{type:"flex"}},[n("i-col",{staticStyle:{color:"white","font-size":"18px"},attrs:{offset:"2",span:"6"}},[n("Icon",{attrs:{type:"social-vimeo"}}),t._v(" "+t._s(t.menus.title))],1),n("i-col",{attrs:{span:"16"}},[n("Row",{attrs:{type:"flex",justify:"end"}},[n("i-col",{attrs:{span:2*t.menus.navs.length}},t._l(t.menus.navs,(function(e,r){return n("Menu-item",{key:r,nativeOn:{click:function(n){return t.show(e.url)}}},[n("Icon",{attrs:{type:e.icon}}),t._v(" "+t._s(e.title)+" ")],1)})),1)],1)],1)],1)],1)],1)],1),n("Row",{attrs:{type:"flex"}},[n("i-col",{attrs:{span:"2"}},[n("Menu",{attrs:{theme:t.theme,"active-key":0}},t._l(t.menus.menus,(function(e){return n("MenuGroup",{key:e,attrs:{title:e.title}},t._l(e.items,(function(e,r){return n("MenuItem",{key:r,nativeOn:{click:function(n){return t.show(e.url)}}},[n("Icon",{attrs:{type:e.icon}}),t._v(" "+t._s(e.title)+" ")],1)})),1)})),1)],1),n("i-col",{attrs:{span:"10"}},[n("iframe",{attrs:{id:"show-iframe",src:t.menus.url,frameborder:"0"}},[t._v(">")])]),n("i-col",{attrs:{span:"2"}})],1)],1)},i=[],a={name:"App",data:function(){return{theme:"light",menus:{title:"hydra帮助手册",url:"/",navs:[{title:"起步"}],menus:[{title:"起步",items:[{title:"概述",url:"/"},{title:"hello world"},{title:"六大服务器"},{title:"启动原理"}]},{title:"构建",items:[{title:"API服务"},{title:"WEB服务"},{title:"WS服务"},{title:"RPC服务"},{title:"CRON服务"},{title:"MQC服务"}]}]}}},mounted:function(){var t=this;this.$fetch("/menus.json").then((function(e){t.menus=e,document.title=e.title})).catch((function(t){console.debug(t)})),this.resetIframe()},methods:{resetIframe:function(){var t=document.getElementById("show-iframe"),e=document.documentElement.clientWidth,n=document.documentElement.clientHeight;t.style.width=e+"px",t.style.height=n+"px",t.attachEvent?t.attachEvent("onload",(function(){var t=document.getElementById("show-iframe").contentWindow.document.getElementsByClassName("markdown-preview");t[0].style.left="35%"})):t.onload=function(){var t=document.getElementById("show-iframe").contentWindow.document.getElementsByClassName("markdown-preview");t[0].style.left="35%"}},show:function(t){this.menus.url=t,console.log("len:",t.length),console.log("sub:",t.substring(0,4)),t.length>4&&"http"==t.substring(0,4)?window.location=t:this.resetIframe()}}},u=a,s=(n("fffb"),n("2877")),l=Object(s["a"])(u,o,i,!1,null,null,null),c=l.exports,f=n("8c4f"),d=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",[t._v(" abcef ")])},p=[],m={name:"componnent1"},h=m,v=Object(s["a"])(h,d,p,!1,null,null,null),y=v.exports;r["default"].use(f["a"]);var w=new f["a"]({mode:"history",routes:[{path:"/",name:"index",component:y}]}),g=n("e069"),b=n.n(g),_=(n("dcad"),n("d3b7"),n("bc3a")),x=n.n(_),O=n("4328");function j(t){var e=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{};return new Promise((function(n,r){x.a.get(t,{params:e}).then((function(t){200==t.status&&n(t.data)})).catch((function(t){r(t)}))}))}function k(t){var e=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{};return e=O.stringify(e),new Promise((function(n,r){x.a.post(t,e).then((function(t){200==t.status&&n(t.data)}),(function(t){r(t)}))}))}x.a.defaults.timeout=5e3,x.a.defaults.withCredentials=!0,x.a.defaults.baseURL="/",x.a.interceptors.request.use((function(t){return t.headers={"Content-Type":"application/x-www-form-urlencoded"},t})),x.a.interceptors.response.use((function(t){return t}),(function(t){return Promise.reject(t)})),r["default"].use(b.a),r["default"].prototype.$post=k,r["default"].prototype.$fetch=j,r["default"].config.productionTip=!1,new r["default"]({router:w,render:function(t){return t(c)}}).$mount("#app")},a210:function(t,e,n){},fffb:function(t,e,n){"use strict";var r=n("a210"),o=n.n(r);o.a}});
//# sourceMappingURL=app.5c7950df.js.map