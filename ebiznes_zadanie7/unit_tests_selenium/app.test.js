const { Builder, By } = require('selenium-webdriver');
require('chromedriver');

(async function() {
  // Inicjalizacja WebDrivera dla Chrome
  const driver = await new Builder().forBrowser('chrome').build();

  try {
    // Otwarcie strony aplikacji
    await driver.get('http://localhost:3000');

    // Test 1: Sprawdź, czy link "Strona główna" istnieje i jest widoczny
    const homeLink = await driver.findElement(By.linkText('Strona główna'));
    const isHomeLinkDisplayed = await homeLink.isDisplayed();
    console.log('Test 1: Strona główna link is displayed:', isHomeLinkDisplayed);

    // Test 2: Sprawdź, czy link "Produkty" istnieje i jest widoczny
    const produktyLink = await driver.findElement(By.linkText('Produkty'));
    const isProduktyLinkDisplayed = await produktyLink.isDisplayed();
    console.log('Test 2: Produkty link is displayed:', isProduktyLinkDisplayed);

    // Test 3: Kliknij link "Produkty" i sprawdź, czy strona Produkty jest wyświetlona
    await produktyLink.click();
    const produktyTitle = await driver.findElement(By.xpath("//h2[contains(text(), 'Produkty')]"));
    const isProduktyDisplayed = await produktyTitle.isDisplayed();
    console.log('Test 3: Produkty page is displayed:', isProduktyDisplayed);

    // Test 4: Sprawdź, czy link "Koszyk" istnieje i jest widoczny
    const koszykLink = await driver.findElement(By.linkText('Koszyk'));
    const isKoszykLinkDisplayed = await koszykLink.isDisplayed();
    console.log('Test 4: Koszyk link is displayed:', isKoszykLinkDisplayed);

    // Test 5: Kliknij link "Koszyk" i sprawdź, czy strona Koszyk jest wyświetlona
    await koszykLink.click();
    const koszykTitle = await driver.findElement(By.xpath("//h2[contains(text(), 'Koszyk')]"));
    const isKoszykDisplayed = await koszykTitle.isDisplayed();
    console.log('Test 5: Koszyk page is displayed:', isKoszykDisplayed);

    // Test 6: Sprawdź, czy link "Płatności" istnieje i jest widoczny
    const platnosciLink = await driver.findElement(By.linkText('Płatności'));
    const isPlatnosciLinkDisplayed = await platnosciLink.isDisplayed();
    console.log('Test 6: Płatności link is displayed:', isPlatnosciLinkDisplayed);

    // Test 7: Kliknij link "Płatności" i sprawdź, czy strona Płatności jest wyświetlona
    await platnosciLink.click();
    const platnosciTitle = await driver.findElement(By.xpath("//h2[contains(text(), 'Płatności')]"));
    const isPlatnosciDisplayed = await platnosciTitle.isDisplayed();
    console.log('Test 7: Płatności page is displayed:', isPlatnosciDisplayed);

    // Test 8: Sprawdź, czy link "Strona główna" nie istnieje
    const nonExistingLink = await driver.findElements(By.linkText('Strona główn'));
    const isNonExistingLinkDisplayed = nonExistingLink.length === 0;
    console.log('Test 8: Strona główna link does not exist:', isNonExistingLinkDisplayed);

    // Test 9: Sprawdź, czy link "Strona główna" istnieje i jest widoczny
    const homeLinkAgain = await driver.findElement(By.linkText('Strona główna'));
    const isHomeLinkAgainDisplayed = await homeLinkAgain.isDisplayed();
    console.log('Test 9: Strona główna link is displayed:', isHomeLinkAgainDisplayed);

    // Test 10: Kliknij link "Strona główna" i sprawdź, czy strona główna jest wyświetlona
    await homeLinkAgain.click();
    const homeTitleAgain = await driver.findElement(By.xpath("//h2[contains(text(), 'Strona główna')]"));
    const isHomeDisplayedAgain = await homeTitleAgain.isDisplayed();
    console.log('Test 10: Home page is displayed:', isHomeDisplayedAgain);
  } finally {
    // Zamknięcie przeglądarki
    await driver.quit();
  }
})();
