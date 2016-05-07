var gulp = require('gulp');
var jasmineBrowser = require('gulp-jasmine-browser');

gulp.task('jasmine', function() {
    return gulp.src(['src/hello_world.html', 'spec/**/*_spec.js'])
        .pipe(jasmineBrowser.specRunner())
        .pipe(jasmineBrowser.server());
});
