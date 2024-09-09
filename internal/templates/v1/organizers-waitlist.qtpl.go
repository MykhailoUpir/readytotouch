// Code generated by qtc from "organizers-waitlist.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package v1

import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

func StreamOrganizersWaitlist(qw422016 *qt422016.Writer, organizerFeature OrganizerFeature, headerProfiles []SocialProviderUser, authQueryParams string, subscribedState bool) {
	qw422016.N().S(`<!DOCTYPE html>
<html lang="en">

<head>
	<title>Yet another anonymous work search</title>
	<meta name="description" content="Yet another anonymous work search">

	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    `)
	streamfavicon(qw422016)
	qw422016.N().S(`
    `)
	streamorganizersFonts(qw422016)
	qw422016.N().S(`
    `)
	streamorganizersWaitlistStyles(qw422016)
	qw422016.N().S(`
	<script src="https://cdn.jsdelivr.net/npm/apexcharts@3.52.0"></script>
</head>

<body class="organizer-inner">
<main class="main-wrapper">
	<header class="header">
  <div class="header__wrapper">
    <a href="#" class="header__logo">
      <img
        width="129"
        height="32"
        class="header__logo-img"
        src="/assets/images/pages/organizer/`)
	qw422016.E().S(organizerFeature.Organizer.Logo)
	qw422016.N().S(`"
        alt="organizer logo"
      />
    </a>
    `)
	var navigation = toFeatureNavigation(organizerFeature.Path)

	qw422016.N().S(`
    <ul class="header__nav">
      <li class="header__nav-item">
        <a href="`)
	qw422016.E().S(navigation.companiesURL)
	qw422016.N().S(`" class="header__nav-link `)
	qw422016.E().S(navigation.companiesActive)
	qw422016.N().S(`">Companies</a>
      </li>
      <li class="header__nav-item">
        <a href="`)
	qw422016.E().S(navigation.vacanciesURL)
	qw422016.N().S(`" class="header__nav-link `)
	qw422016.E().S(navigation.vacanciesActive)
	qw422016.N().S(`">Vacancies</a>
      </li>
    </ul>
    `)
	streamorganizersHeaderStars(qw422016)
	qw422016.N().S(`
    `)
	if len(headerProfiles) > 0 {
		qw422016.N().S(`
    `)
		streamorganizersHeaderProfile(qw422016, headerProfiles)
		qw422016.N().S(`
    `)
	} else {
		qw422016.N().S(`
    <a href="/organizers/`)
		qw422016.E().S(organizerFeature.Organizer.Alias)
		qw422016.N().S(`/welcome`)
		qw422016.E().S(authQueryParams)
		qw422016.N().S(`" class="button button--small-padding button--black header__login-button">Log in</a>
    `)
	}
	qw422016.N().S(`
  </div>
</header>

<div class="container">
  <nav aria-label="breadcrumb" aria-labelledby="navigation through the bread crumbs" class="breadcrumb">
    <ul class="breadcrumb__list">
      <li class="breadcrumb__item">
        <a class="breadcrumb__link" href="/">Main</a>
      </li>
      <li class="breadcrumb__item">
        <a class="breadcrumb__link" href="/organizers">Organizers</a>
      </li>
      <li class="breadcrumb__item">
        <a class="breadcrumb__link" href="javascript:void(0);">`)
	qw422016.E().S(organizerFeature.Organizer.Title)
	qw422016.N().S(`</a>
      </li>
      <li class="breadcrumb__item">
        <span class="breadcrumb__page" aria-current="page">`)
	qw422016.E().S(organizerFeature.Title)
	qw422016.N().S(`</span>
      </li>
    </ul>
  </nav>
</div>

<section class="hero">
  <div class="container hero__container">
    <div class="hero__overlay">
      <img
        class="hero__image"
        width="56"
        height="67"
        srcset="/assets/images/pages/organizer/eyes@2x.png 2x"
        src="/assets/images/pages/organizer/eyes.png"
        alt="eyes icon"
      />
      <h1 class="hero__headline">You've come across something we're still working on</h1>
      `)
	var (
		subscribeStateClass   = ""
		unsubscribeStateClass = ""
	)

	if subscribedState {
		subscribeStateClass = "d-none"
	} else {
		unsubscribeStateClass = "d-none"
	}

	qw422016.N().S(`
      <p class="hero__description hero__description--with-icon js-feature-unsubscribe `)
	qw422016.E().S(unsubscribeStateClass)
	qw422016.N().S(`">
        <img
          class="hero__description-image"
          width="20"
          height="20"
          srcset="/assets/images/pages/organizer/check-mark@2x.png 2x"
          src="/assets/images/pages/organizer/check-mark.png"
          alt="check mar icon"
        />
        You are subscribed to be notified when this page is available
      </p>
      <a href="javascript:void(0);" class="button button--small-padding button--bordered-gray hero__button js-feature-unsubscribe `)
	qw422016.E().S(unsubscribeStateClass)
	qw422016.N().S(`">Unsubscribe</a>

      <p class="hero__description js-feature-subscribe `)
	qw422016.E().S(subscribeStateClass)
	qw422016.N().S(`">Subscribe to be notified when this page is available</p>
      <a href="javascript:void(0);" class="button button--small-padding button--black hero__button js-feature-subscribe `)
	qw422016.E().S(subscribeStateClass)
	qw422016.N().S(`">Subscribe</a>
    </div>
  </div>
</section>

<section class="project-statistics">
  <div class="container project-statistics__container">
    <div class="stats">
      <div class="stats__item">
        <div class="stats__header">
          <h3 class="stats__title">Page views</h3>
        </div>
        <div class="stats__counters">
          <div class="stats__counters-item">
            <p class="stats__counters-views">Total views</p>
            <p class="stats__counters-item-number js-stats-total-views">∞</p>
          </div>
          <div class="stats__counters-item">
            <p class="stats__counters-views">Last month's views</p>
            <p class="stats__counters-item-number js-stats-last-month-views">∞</p>
          </div>
        </div>
        <div class="stats__chart stats__chart--page-views-statistics js-chart-views-statistics"></div>
      </div>
      <div class="stats__item">
        <div class="stats__header">
          <h3 class="stats__title">Subscribers</h3>
        </div>
        <div class="stats__counters">
          <div class="stats__counters-item">
            <p class="stats__counters-views">Total subscribers</p>
            <p class="stats__counters-item-number js-stats-total-subscribers">∞</p>
          </div>
          <div class="stats__counters-item">
            <p class="stats__counters-views">Last month’s subscribers</p>
            <p class="stats__counters-item-number js-stats-last-month-subscribers">∞</p>
          </div>
        </div>
        <div class="stats__chart">
          <div class="stats__chart stats__chart--subscribers-statistics js-chart-subscribers-statistics"></div>
        </div>
      </div>
    </div>
  </div>
</section>


</main>
<footer class="footer">
  <div class="footer__overlay">
    <div class="footer__wrapper">
      <div class="footer__group">
        <div class="footer__info">
          <a href="/" class="footer__title">
            <h3 class="footer__title-link">ReadyToTouch</h3>
          </a>
          <p class="footer__subtitle">Space to help you find jobs</p>
          <img width="98"
               height="64"
               loading="lazy"
               class="footer__ukraine-map"
               src="/assets/images/pages/online-new/ukraine-map.svg"
               alt="Map of Ukraine"
          >
        </div>
        <div class="footer__right-side">
          <iframe class="footer__stars"
                  src="https://ghbtns.com/github-btn.html?user=readytotouch&repo=readytotouch&type=star&count=true&size=large"
                  width="168"
                  height="33"
                  title="GitHub"
          ></iframe>
          <div class="footer__support">
            <a href="https://savelife.in.ua/en/" target="_blank" class="button button--bordered-gray footer__link">
              <figure class="footer__support-figure">
                <img width="49"
                     height="23"
                     src="/assets/images/pages/common-images/come_back_alive.svg"
                     alt="come back alive logo"
                >
                <figcaption class="footer__support-caption">Support</figcaption>
              </figure>
              <div class="footer__support-icon-box">
                <img width="18"
                     height="18"
                     src="/assets/images/pages/common/external-link.svg"
                     alt="external link icon"
                >
              </div>
            </a>
            <a href="https://war.ukraine.ua/" target="_blank" class="button button--bordered-gray footer__link">
              war.ukraine.ua
              <div class="footer__support-icon-box">
                <img width="18"
                     height="18"
                     src="/assets/images/pages/common/external-link.svg"
                     alt="external link icon"
                >
              </div>
            </a>
          </div>
        </div>
      </div>
      <div class="footer__copyrights">
        &copy; 2024 Yaroslav Podorvanov
        <img width="24"
             height="24"
             class="footer__flag-ua"
             src="/assets/images/pages/online-new/ua_flag_with_waves.svg"
             alt="Flag of Ukraine"
        >
      </div>
    </div>
  </div>
</footer>
<script src="/assets/js/waitlist-stats-app.js?`)
	qw422016.N().D(appVersion)
	qw422016.N().S(`"></script>
</body>
</html>
`)
}

func WriteOrganizersWaitlist(qq422016 qtio422016.Writer, organizerFeature OrganizerFeature, headerProfiles []SocialProviderUser, authQueryParams string, subscribedState bool) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamOrganizersWaitlist(qw422016, organizerFeature, headerProfiles, authQueryParams, subscribedState)
	qt422016.ReleaseWriter(qw422016)
}

func OrganizersWaitlist(organizerFeature OrganizerFeature, headerProfiles []SocialProviderUser, authQueryParams string, subscribedState bool) string {
	qb422016 := qt422016.AcquireByteBuffer()
	WriteOrganizersWaitlist(qb422016, organizerFeature, headerProfiles, authQueryParams, subscribedState)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}