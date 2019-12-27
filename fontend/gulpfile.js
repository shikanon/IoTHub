// gulpfile.js
var path = require('path')
var gulp = require('gulp')
var cleanCSS = require('gulp-clean-css')
var cssWrap = require('gulp-css-wrap')
gulp.task('css-wrap', function () {
    return gulp.src(path.resolve('./theme/index.css'))
    /* 找需要添加命名空间的css文件，支持正则表达式 */
        .pipe(cssWrap({
            selector: '.custom-02' /* 添加的命名空间 */
        }))
        .pipe(cleanCSS())
        .pipe(gulp.dest('src/assets/css/theme/02')) /* 存放的目录 */
})


/****
 * 生成主题步骤
 * 1.修改element-variables.scss
 * 2.编译，在控制台输入命令 et ，编译生成theme文件夹
 * 3.在main.js中引入 import '../theme/index.css'
 * 4.修改gulpfile.js，执行gulp输出  gulp css-wrap ，生成路径src/assets/css/theme/10,此时文件夹下没有fonts文件夹，把生成主题时生成的fonts复制到此处即可
 * 5.在main.js中引入生成的主题文件，import './assets/css/theme/10/index.css'
 */