//Magus func to spam submit on CvZ
// tu fais F12 et tu colles ça dans la fenetre de ton navigateur quand t'es sur l'IDE
// ça va submit en boucle ton code

(function() {
  var count = 0,
      max = 0,
      fail = 0,
      average = 0;

  $('.cg-ide-testcases .testcase-button').first().click();
      
  window.magusSubmitBot = setInterval(function() {
    console.log('Counter :', ++count);

    $('.cg-ide-actions button.submit').click();

    setTimeout(function() {
      $('.mask.angular-animate.ng-scope').click();

      var score = parseInt($('[ng-bind="report.criteriaScore"]').html(), 10);
      console.log('Score :', score);

      if (score > max) {
        max = score;
      }

      average = average*(count - 1)/count + score/count;

      if ($('.cg-report-score-progress-value').html() != '100') {
        fail += 1;
        console.log('Fail !');
      }

      console.log('Max :', max);
      console.log('Average :', Math.round(average));
      console.log('Fails : ', 100*(fail / count) + '%');
    }, 18000);
  }, 180000);
})();
