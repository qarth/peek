import { chromium } from 'playwright';

async function scrapeSeekMiningEngineers() {
  const browser = await chromium.launch();
  const page = await browser.newPage();
  const url = 'https://www.seek.com.au/mining-engineer-jobs/in-All-Australia';
  await page.goto(url, { waitUntil: 'domcontentloaded' });

  // Wait for job cards to load
  await page.waitForSelector("article[data-automation='normalJob']");

  const jobs = await page.$$eval("article[data-automation='normalJob']", (articles) => {
    return articles.map((article) => {
      const titleEl = article.querySelector("a[data-automation='jobTitle']");
      const companyEl = article.querySelector("a[data-automation='jobCompany']");
      const locationEl = article.querySelector("a[data-automation='jobLocation']");
      return {
        title: titleEl?.textContent?.trim() || '',
        link: titleEl ? 'https://www.seek.com.au' + titleEl.getAttribute('href') : '',
        company: companyEl?.textContent?.trim() || '',
        location: locationEl?.textContent?.trim() || '',
      };
    });
  });

  jobs.forEach((job, i) => {
    console.log(`Job ${i + 1}:`);
    console.log(`  Title: ${job.title}`);
    console.log(`  Link: ${job.link}`);
    console.log(`  Company: ${job.company}`);
    console.log(`  Location: ${job.location}`);
    console.log();
  });

  await browser.close();
}

scrapeSeekMiningEngineers();
