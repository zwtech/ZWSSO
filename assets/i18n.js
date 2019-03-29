(function(k){function t(a){a.debug&&(g("callbackIfComplete()"),g("totalFiles: "+a.totalFiles),g("filesLoaded: "+a.filesLoaded));a.async&&a.filesLoaded===a.totalFiles&&a.callback&&a.callback()}function n(a,d){d.debug&&g("loadAndParseFiles");null!==a&&0<a.length?u(a[0],d,function(){a.shift();n(a,d)}):t(d)}function u(a,d,b){d.debug&&(g("loadAndParseFile('"+a+"')"),g("totalFiles: "+d.totalFiles),g("filesLoaded: "+d.filesLoaded));null!==a&&"undefined"!==typeof a&&k.ajax({url:a,async:d.async,cache:d.cache,
    dataType:"text",success:function(h,c){d.debug&&(g("Succeeded in downloading "+a+"."),g(h));v(h,d);b()},error:function(h,c,e){d.debug&&g("Failed to download or parse "+a+". errorThrown: "+e);404===h.status&&--d.totalFiles;b()}})}function v(a,d){var b="",h=a.split(/\n/),c=/(\{\d+})/g,e=/\{(\d+)}/g,g=/(\\u.{4})/ig;h.forEach(function(a,w){a=a.trim();if(0<a.length&&"#"!=a.match("^#")){var l=a.split("=");if(0<l.length){for(var m=decodeURI(l[0]).trim(),f=1==l.length?"":l[1];"\\"===f.match(/\\$/);)f=f.substring(0,
    f.length-1),f+=h[++w].trimRight();for(var p=2;p<l.length;p++)f+="="+l[p];f=f.trim();if("map"==d.mode||"both"==d.mode)(l=f.match(g))&&l.forEach(function(a){f=f.replace(a,x(a))}),k.i18n.map[m]=f;if("vars"==d.mode||"both"==d.mode)if(f=f.replace(/"/g,'\\"'),y(m),c.test(f)){var n=!0,q="",r=[];f.split(c).forEach(function(a){!c.test(a)||0!==r.length&&-1!=r.indexOf(a)||(n||(q+=","),q+=a.replace(e,"v$1"),r.push(a),n=!1)});b+=m+"=function("+q+"){";m='"'+f.replace(e,'"+v$1+"')+'"';b+="return "+m+";};"}else b+=
    m+'="'+f+'";'}}});eval(b);d.filesLoaded+=1}function y(a){if(/\./.test(a)){var d="";a.split(/\./).forEach(function(a,h){0<h&&(d+=".");d+=a;eval("typeof "+d+' == "undefined"')&&eval(d+"={};")})}}function x(a){var d=[];a=parseInt(a.substr(2),16);0<=a&&a<Math.pow(2,16)&&d.push(a);return d.reduce(function(a,d){return a+String.fromCharCode(d)},"")}k.i18n={};k.i18n.map={};var g=function(a){window.console&&console.log("i18n::"+a)};k.i18n.properties=function(a){a=k.extend({name:"Messages",language:"",path:"",
    mode:"vars",cache:!1,debug:!1,encoding:"UTF-8",async:!1,callback:null},a);a.path.match(/\/$/)||(a.path+="/");a.language=this.normaliseLanguageCode(a.language);var d=a.name&&a.name.constructor===Array?a.name:[a.name];a.totalFiles=2*d.length+(5<=a.language.length?d.length:0);a.debug&&g("totalFiles: "+a.totalFiles);a.filesLoaded=0;d.forEach(function(b){var d=a.path+b+".properties";var c=a.language.substring(0,2);c=a.path+b+"_"+c+".properties";if(5<=a.language.length){var e=a.language.substring(0,5);
    b=a.path+b+"_"+e+".properties";d=[d,c,b]}else d=[d,c];n(d,a)});a.callback&&!a.async&&a.callback()};k.i18n.prop=function(a){var d,b=k.i18n.map[a];if(null===b)return"["+a+"]";var h;2==arguments.length&&k.isArray(arguments[1])&&(h=arguments[1]);var c;if("string"==typeof b){for(c=0;-1!=(c=b.indexOf("\\",c));)b="t"==b.charAt(c+1)?b.substring(0,c)+"\t"+b.substring(c++ +2):"r"==b.charAt(c+1)?b.substring(0,c)+"\r"+b.substring(c++ +2):"n"==b.charAt(c+1)?b.substring(0,c)+"\n"+b.substring(c++ +2):"f"==b.charAt(c+
    1)?b.substring(0,c)+"\f"+b.substring(c++ +2):"\\"==b.charAt(c+1)?b.substring(0,c)+"\\"+b.substring(c++ +2):b.substring(0,c)+b.substring(c+1);var e=[];for(c=0;c<b.length;)if("'"==b.charAt(c))if(c==b.length-1)b=b.substring(0,c);else if("'"==b.charAt(c+1))b=b.substring(0,c)+b.substring(++c);else{for(d=c+2;-1!=(d=b.indexOf("'",d));)if(d==b.length-1||"'"!=b.charAt(d+1)){b=b.substring(0,c)+b.substring(c+1,d)+b.substring(d+1);c=d-1;break}else b=b.substring(0,d)+b.substring(++d);-1==d&&(b=b.substring(0,c)+
    b.substring(c+1))}else if("{"==b.charAt(c))if(d=b.indexOf("}",c+1),-1==d)c++;else{var g=parseInt(b.substring(c+1,d));!isNaN(g)&&0<=g?(c=b.substring(0,c),""!==c&&e.push(c),e.push(g),c=0,b=b.substring(d+1)):c=d+1}else c++;""!==b&&e.push(b);b=e;k.i18n.map[a]=e}if(0===b.length)return"";if(1==b.length&&"string"==typeof b[0])return b[0];e="";c=0;for(d=b.length;c<d;c++)e="string"==typeof b[c]?e+b[c]:h&&b[c]<h.length?e+h[b[c]]:!h&&b[c]+1<arguments.length?e+arguments[b[c]+1]:e+("{"+b[c]+"}");return e};k.i18n.normaliseLanguageCode=
    function(a){if(!a||2>a.length)a=navigator.languages&&0<navigator.languages.length?navigator.languages[0]:navigator.language||navigator.userLanguage||"en";a=a.toLowerCase();a=a.replace(/-/,"_");3<a.length&&(a=a.substring(0,3)+a.substring(3).toUpperCase());return a}})(jQuery);
