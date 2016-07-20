var gulp = require('gulp');
var jasmineBrowser = require('gulp-jasmine-browser');
var webpack = require('gulp-webpack');

gulp.task('jasmine', function() {
    return gulp.src(['src/task_list.html', 'spec/**/*_spec.js'])
        .pipe(jasmineBrowser.specRunner())
        .pipe(jasmineBrowser.server());
});

gulp.task('build', function() {
    return gulp.src('src/index.js')
        .pipe(webpack({
            watch: true,
            module: {
                preLoaders: [
                    { test: /\/src\/.+\.js$/, loader: 'polymer-loader?templateExtension=html' }
                ],
                loaders: [
                    { test: /\.html$/, loader: "html" }
                ]
            },
            output: {
                filename: "bundle.js"
            }
        }))
        .pipe(gulp.dest('../build/public'));
});