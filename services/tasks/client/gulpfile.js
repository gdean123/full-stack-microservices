var gulp = require('gulp');
var jasmineBrowser = require('gulp-jasmine-browser');

gulp.task('jasmine', function() {
    return gulp.src(['src/task_list.html', 'spec/**/*_spec.js'])
        .pipe(jasmineBrowser.specRunner())
        .pipe(jasmineBrowser.server());
});
