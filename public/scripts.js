const inputUrl = document.querySelector('#input-url');
const outputUrl = document.querySelector('.sm-target-link-container');
const smOutputUrl = document.querySelector('.sm-target-link');
const targetUrlElement = document.querySelector('.target-link');
const targetUrl = document.querySelector('.copy-icon-box');
const smTargetUrl = document.querySelector('.sm-copy-icon-box');
const clear = document.querySelector('.close-icon');
const errorEl = document.querySelector('.error');
const submitBtn = document.getElementById('btn-submit');
const mobileDeviceWidth = 600;
let deviceWidth = window.innerWidth;

const shortenUrl = async () => {
  try {
    const response = await fetch('/add', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        target: inputUrl.value,
      }),
    });
    const res = await response.json();
    return res.slug;
  } catch (error) {}
};

const copyLink = () => {
  const targetUrlText = targetUrlElement.textContent;

  if (targetUrlText.length > 0) {
    navigator.clipboard.writeText(targetUrlText);
  }
};

const adjustElementVisibilityByScreenWidth = () => {
  deviceWidth <= mobileDeviceWidth
    ? (outputUrl.style.display = 'block')
    : (outputUrl.style.display = 'none');
};

const handleShortenButtonClick = async () => {
  const inputUrlValue = inputUrl.value;

  if (isValidHttpUrl(inputUrlValue)) {
    try {
      const shortenedUrl = await shortenUrl();
      const targetUrl = `${window.location.origin}/${shortenedUrl}`;

      errorEl.style.display = 'none';

      adjustElementVisibilityByScreenWidth();

      targetUrlElement.textContent = targetUrl;
      smOutputUrl.textContent = targetUrl;
    } catch {}
  } else {
    errorEl.style.display = 'block';
  }
};

const clearLink = () => {
  inputUrl.value = '';
};

const isValidHttpUrl = (string) => {
  let url;
  try {
    url = new URL(string);
  } catch (_) {
    return false;
  }

  return url.protocol === 'http:' || url.protocol === 'https:';
};

submitBtn.addEventListener('click', handleShortenButtonClick);
targetUrl.addEventListener('click', copyLink);
smTargetUrl.addEventListener('click', copyLink);
clear.addEventListener('click', clearLink);

window.addEventListener('resize', (e) => {
  deviceWidth = window.innerWidth;
  const isMobile = deviceWidth <= mobileDeviceWidth;

  outputUrl.style.display =
    isMobile && smOutputUrl.textContent.length > 0 ? 'block' : 'none';
});
