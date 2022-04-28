package main

import "testing"

var htmlSample []byte = []byte(`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" lang="en">
    <body id="ctl00_CWBody">
        <div id="inova_alert_div">
<div class="inova_alert_container">
<div class="inova_alert_message">
<h2>
<span>COVID-19 Information for the Community</span>
</h2>
<div _rdeditor_temp="1"> Find information about updates, cancellations and how Inova is taking the appropriate steps to protect the safety of our caregivers, patients, team members and the community.</div>
</div>
<a href="https://www.inova.org/covid19" target="_blank" class="inova_alert_btn" aria-label="COVID-19 Information for the
Community"> See the Latest</a>
<div class="inova_alert_clear">&nbsp;</div>
</div>
</div>

        <!-- Google Tag Manager -->


<noscript>
    <iframe src="//www.googletagmanager.com/ns.html?id=GTM-KHWC54"
        height="0" width="0" style="display: none; visibility: hidden"></iframe>
</noscript>
<script>(function (w, d, s, l, i) {
w[l] = w[l] || []; w[l].push({
'gtm.start':
new Date().getTime(), event: 'gtm.js'
}); var f = d.getElementsByTagName(s)[0],
j = d.createElement(s), dl = l != 'dataLayer' ? '&l=' + l : ''; j.async = true; j.src =
'//www.googletagmanager.com/gtm.js?id=' + i + dl; f.parentNode.insertBefore(j, f);
})(window, document, 'script', 'dataLayer', 'GTM-KHWC54');</script>

        <!-- End Google Tag Manager -->


<div id="fb-root"></div>

<script>(function (d, s, id) {
    var js, fjs = d.getElementsByTagName(s)[0];
    if (d.getElementById(id)) return;
    js = d.createElement(s); js.id = id;
    js.src = "//connect.facebook.net/en_US/all.js#xfbml=1";
    fjs.parentNode.insertBefore(js, fjs);
}(document, 'script', 'facebook-jssdk'));</script>

<!-- imagemap for browser update notification-->
<map id="browserUpdateMap" name="browserUpdateMap">
    <area alt="Close" title="Close" shape="circle" coords="420,16,8" style="cursor:pointer;" onclick="tb_remove();"/>
    <area alt="Get Latest Internet Explorer" title="Get Latest Internet Explorer" shape="rect" coords="51,58,156,144" href="http://www.microsoft.com/IE9" target="_blank"/>
    <area alt="Get Latest Firefox" title="Get Latest Firefox" shape="rect" coords="167,58,272,144" href="http://www.mozilla.com/firefox/fx/" target="_blank"/>
    <area alt="Get Latest Chrome" title="Get Latest Chrome" shape="rect" coords="248,58,389,144" href="http://www.google.com/chrome" target="_blank"/>
</map>


        <form method="post" action="/emergency-room-wait-times/default.aspx" id="aspnetForm">
<div class="aspNetHidden">
<input type="hidden" name="ctl00_ctl05_TSM" id="ctl00_ctl05_TSM" value="" />
<input type="hidden" name="ctl00_ctl06_TSSM" id="ctl00_ctl06_TSSM" value="" />
<input type="hidden" name="__EVENTTARGET" id="__EVENTTARGET" value="" />
<input type="hidden" name="__EVENTARGUMENT" id="__EVENTARGUMENT" value="" />
<input type="hidden" name="__VIEWSTATE" id="__VIEWSTATE" value="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" />
</div>

<script type="text/javascript">
//<![CDATA[
var theForm = document.forms['aspnetForm'];
function __doPostBack(eventTarget, eventArgument) {
    if (!theForm.onsubmit || (theForm.onsubmit() != false)) {
        theForm.__EVENTTARGET.value = eventTarget;
        theForm.__EVENTARGUMENT.value = eventArgument;
        theForm.submit();
    }
}
//]]>
</script>


<script src="/WebResource.axd?d=Da-y10GH5RwP9s_36ZNvCXkj34Y_yW71zne3WKrFyTcXl0lI_vMbHoC76Pr-SLc3Mvoq_F2ofM0XFEnpuIbRnOkyb981&amp;t=637811551220000000" type="text/javascript"></script>


<script src="/Telerik.Web.UI.WebResource.axd?_TSM_HiddenField_=ctl00_ctl05_TSM&amp;compress=1&amp;_TSM_CombinedScripts_=%3b%3bSystem.Web.Extensions%2c+Version%3d4.0.0.0%2c+Culture%3dneutral%2c+PublicKeyToken%3d31bf3856ad364e35%3aen-US%3a5bc44d53-7cae-4d56-af98-205692fecf1f%3aea597d4b%3ab25378d2" type="text/javascript"></script>
<script src="https://d2i2wahzwrm1n5.cloudfront.net/ajaxz/2014.3.1209/Common/Core.js" type="text/javascript"></script>
<script src="https://d2i2wahzwrm1n5.cloudfront.net/ajaxz/2014.3.1209/Common/jQuery.js" type="text/javascript"></script>
<script src="https://d2i2wahzwrm1n5.cloudfront.net/ajaxz/2014.3.1209/Common/jQueryPlugins.js" type="text/javascript"></script>
<script src="https://d2i2wahzwrm1n5.cloudfront.net/ajaxz/2014.3.1209/Common/Scrolling/ScrollingScripts.js" type="text/javascript"></script>
<script src="https://d2i2wahzwrm1n5.cloudfront.net/ajaxz/2014.3.1209/Common/Navigation/OData/OData.js" type="text/javascript"></script>
<script src="https://d2i2wahzwrm1n5.cloudfront.net/ajaxz/2014.3.1209/Common/AnimationFramework/AnimationFramework.js" type="text/javascript"></script>
<script src="https://d2i2wahzwrm1n5.cloudfront.net/ajaxz/2014.3.1209/Common/Navigation/NavigationScripts.js" type="text/javascript"></script>
<script src="https://d2i2wahzwrm1n5.cloudfront.net/ajaxz/2014.3.1209/Common/Navigation/OverlayScript.js" type="text/javascript"></script>
<script src="https://d2i2wahzwrm1n5.cloudfront.net/ajaxz/2014.3.1209/Menu/RadMenuScripts.js" type="text/javascript"></script>
<script src="https://d2i2wahzwrm1n5.cloudfront.net/ajaxz/2014.3.1209/Menu/MenuItem/RadMenuItem.js" type="text/javascript"></script>
<script src="https://d2i2wahzwrm1n5.cloudfront.net/ajaxz/2014.3.1209/Menu/Views/ClassicView.js" type="text/javascript"></script>
<div class="aspNetHidden">

	<input type="hidden" name="__VIEWSTATEGENERATOR" id="__VIEWSTATEGENERATOR" value="76BCB8B4" />
	<input type="hidden" name="__EVENTVALIDATION" id="__EVENTVALIDATION" value="/wEdAAM6KWKYkIlOwc6bQjXoEd9qGU7NIF7bBO3k2pNp+QiRKobCAA4tlP3aVATD1lrDrXmGg5eaGZRwafyae23b4CIbuLkEHQ==" />
</div>
            <script type="text/javascript">
//<![CDATA[
Sys.WebForms.PageRequestManager._initialize('ctl00$ctl05', 'aspnetForm', [], [], [], 90, 'ctl00');
//]]>
</script>

            <a id="top" name="top"></a>

            <script type="text/javascript">
                $(window).on('load', function () {
                    $(".searchInput").focus(function () {
                        if ($(this).val() == "search...")
                            $(this).val("");
                    });

                    $(".searchInput").blur(function () {
                        if ($(this).val() == "")
                            $(this).val("search...");
                    });
                });
            </script>
            <div class="frameWrap">
                <div class="header">
                    <div id="logo">
                        <a href="/?sid=22">
                            <img src="/images/design010/logo.png" alt=""  height="56" /></a>
                    </div>
                    <div id="headerRight">
                        <div id="headerMenu">
                            <div id="utilityWrap">
                                <span id="sharing">
                                    <script type="text/javascript">
                                        /* <![CDATA[ */
                                        if (typeof (SHARETHIS) !== 'undefined' && SHARETHIS != null) {
                                            SHARETHIS.addEntry({
                                                title: 'Share This'
                                            }, { offsetLeft: -385, offsetTop: -25 });
                                        }
                                        /* ]]> */
                                    </script>
                                </span>
                                / <a href="javascript:void(printSpecial('printReady','010'))" id="print">Print</a>
                                / <span class="textSize">Text Size:</span>
                                <a href="javascript:ts('printReady',1)" class="Abig">A</a>
                                <a href="javascript:ts('printReady',0)" class="Amedium">A</a>
                                <a href="javascript:ts('printReady',-1)" class="Asmall">A</a>


											<div id="data-reactroot" style="float:right;" class="desktopOverlay"></div>


                            </div>

                            <div id="searchWrap">
                                <input name="ctl00$txtSearch010" type="text" id="ctl00_txtSearch010" value="search..." class="searchInput" />
                                <input type="image" name="ctl00$btnSearch010" id="ctl00_btnSearch010" class="searchBtn" src="/images/design010/searchBtn.jpg" />
                            </div>

                        </div>
                    </div>
                </div>

                <div class="navBar">
                    <div id="ctl00_PrimaryNavigation" class="RadMenu RadMenu_Design010 rmSized" style="width:100%;">
	<!-- 2014.3.1209.40 --><ul class="rmRootGroup rmHorizontal">
		<li class="rmItem rmFirst"><a class="rmLink rmRootLink" href="/our-spine-team"><span class="rmText">Our Physicians</span></a><div class="rmSlide">
			<ul class="rmVertical rmGroup rmLevel1">
				<li class="rmItem rmFirst"><a class="rmLink" href="/our-spine-team/medical-directors/"><span class="rmText">Medical Directors</span></a></li><li class="rmItem "><a class="rmLink" href="/our-spine-team/physicians"><span class="rmText">Member Physicians</span></a></li><li class="rmItem rmLast"><a class="rmLink" href="/our-spine-team/care-coordination-team"><span class="rmText">Spine Care Coordination Team</span></a></li>
			</ul>
		</div></li><li class="rmItem rmSeparator"><span class="rmText"></span></li><li class="rmItem "><a class="rmLink rmRootLink" href="/spine-conditions"><span class="rmText">Conditions</span></a><div class="rmSlide">
			<ul class="rmVertical rmGroup rmLevel1">
				<li class="rmItem rmFirst"><a class="rmLink" href="/spine-conditions/types-of-conditions"><span class="rmText">Types of Spine Conditions</span></a></li><li class="rmItem "><a class="rmLink" href="/spine-conditions/symptoms"><span class="rmText">Symptoms of Spine Conditions</span></a></li><li class="rmItem "><a class="rmLink" href="/spine-conditions/diagnosing"><span class="rmText">Diagnosing Spine and Back Conditions</span></a></li><li class="rmItem rmLast"><a class="rmLink" href="/spine-conditions-resources/"><span class="rmText">Resource Center</span></a></li>
			</ul>
		</div></li><li class="rmItem rmSeparator"><span class="rmText"></span></li><li class="rmItem "><a class="rmLink rmRootLink" href="/treatment-options/"><span class="rmText">Treatment</span></a><div class="rmSlide">
			<ul class="rmVertical rmGroup rmLevel1">
				<li class="rmItem rmFirst"><a class="rmLink" href="/treatment-options/surgery"><span class="rmText">Surgical Procedures</span></a><div class="rmSlide">
					<ul class="rmVertical rmGroup rmLevel2">
						<li class="rmItem rmFirst rmLast"><a class="rmLink" href="/treatment-options/surgery/minimally-invasive-spine-surgery"><span class="rmText">Minimally Invasive Spine Surgery</span></a></li>
					</ul>
				</div></li><li class="rmItem "><a class="rmLink" href="/treatment-options/physical-therapy"><span class="rmText">Conservative Care and Rehabilitation</span></a></li><li class="rmItem rmLast"><a class="rmLink" href="/treatment-options/pain-management"><span class="rmText">Pain Management Procedures</span></a></li>
			</ul>
		</div></li><li class="rmItem rmSeparator"><span class="rmText"></span></li><li class="rmItem "><a class="rmLink rmRootLink" href="/resource-center/prep-for-surgery"><span class="rmText">Resource Center</span></a><div class="rmSlide">
			<ul class="rmVertical rmGroup rmLevel1">
				<li class="rmItem rmFirst"><a class="rmLink" href="/resource-center/blog-posts"><span class="rmText">Blog Posts from Our Doctors</span></a></li><li class="rmItem "><a class="rmLink" href="/patient-resources/preparing-for-surgery-or-treatment"><span class="rmText">Preparing for Your Surgery or Treatment</span></a></li><li class="rmItem "><a class="rmLink" href="/spine-doctor-videos"><span class="rmText">Spine Doctor Videos</span></a></li><li class="rmItem "><a class="rmLink" href="/resource-center/forms"><span class="rmText">Patient Forms</span></a></li><li class="rmItem "><a class="rmLink" href="/resource-center/online-back-in-action"><span class="rmText">Online Back in Action Class </span></a></li><li class="rmItem rmLast"><a class="rmLink" href="/resource-center/health-library"><span class="rmText">Health Library Resources</span></a></li>
			</ul>
		</div></li><li class="rmItem rmSeparator"><span class="rmText"></span></li><li class="rmItem "><a class="rmLink rmRootLink" href="/virtual-visits"><span class="rmText">Virtual Visits</span></a></li><li class="rmItem rmSeparator"><span class="rmText"></span></li><li class="rmItem "><a class="rmLink rmRootLink" href="/locations"><span class="rmText">Locations</span></a><div class="rmSlide">
			<ul class="rmVertical rmGroup rmLevel1">
				<li class="rmItem rmFirst"><a class="rmLink" href="/locations/inova-spine-assessment"><span class="rmText">Inova Spine Assessment Centers</span></a></li><li class="rmItem "><a class="rmLink" href="/locations/inova-alexandria-hospital"><span class="rmText">Inova Alexandria Hospital</span></a></li><li class="rmItem "><a class="rmLink" href="/locations/inova-fairfax-hospital"><span class="rmText">Inova Fairfax Hospital</span></a></li><li class="rmItem "><a class="rmLink" href="/locations/inova-fair-oaks-hospital"><span class="rmText">Inova Fair Oaks Hospital</span></a></li><li class="rmItem "><a class="rmLink" href="/locations/inova-loudoun-hospital"><span class="rmText">Inova Loudoun Hospital</span></a></li><li class="rmItem rmLast"><a class="rmLink" href="/locations/inova-mount-vernon-hospital"><span class="rmText">Inova Mount Vernon Hospital</span></a></li>
			</ul>
		</div></li><li class="rmItem rmSeparator"><span class="rmText"></span></li><li class="rmItem rmLast"><a class="rmLink rmRootLink" href="/about-spine/index"><span class="rmText">About Us</span></a></li>
	</ul><input id="ctl00_PrimaryNavigation_ClientState" name="ctl00_PrimaryNavigation_ClientState" type="hidden" />
</div>
                </div>
<div id="mobile-nav">
				<ul id="mobile-menu">
					<li><a href="/our-spine-team">Our Physicians</a>
						<ul class="child-nav">
							<li><a href="/our-spine-team/medical-directors/">Medical Directors</a></li>
							<li><a href="/our-spine-team/physicians">Member Physicians</a></li>
							<li><a href="/our-spine-team/care-coordination-team">Spine Care Coordination Team</a></li>
						</ul>
					</li>
					<li><a href="/spine-conditions">Spine Conditions</a>
						<ul class="child-nav">
							<li><a class="rmLink" href="/spine-conditions/types-of-conditions">Types of Spine Conditions</a></li>
							<li><a class="rmLink" href="/spine-conditions/symptoms">Symptoms of Spine Conditions</a></li>
							<li><a class="rmLink" href="/spine-conditions/diagnosing">Diagnosing Spine and Back Conditions</a></li>
							<li><a class="rmLink" href="/spine-conditions-resources/">Resource Center</a></li>
						</ul>
					</li>
					<li><a href="/treatment-options/">Treatment Options</a>
						<ul class="child-nav">
							<li><a class="rmLink" href="/treatment-options/surgery">Surgical Procedures</a></li>
							<li><a class="rmLink" href="/treatment-options/physical-therapy">Conservative Care and Rehabilitation</a></li>
							<li><a class="rmLink" href="/treatment-options/pain-management">Pain Management Procedures</a></li>
						</ul>
					</li>
					<li><a href="/resource-center/prep-for-surgery">Resource Center</a>
						<ul class="child-nav">
							<li><a class="rmLink" href="/patient-resources/preparing-for-surgery-or-treatment">Preparing for Your Surgery or Treatment</a></li>
							<li><a class="rmLink" href="/spine-doctor-videos">Spine Doctor Videos</a></li>
							<li><a class="rmLink" href="/resource-center/forms">Inova Spine Patient Forms</a></li>
							<li><a class="rmLink" href="/resource-center/online-back-in-action">Online Back in Action Class </a></li>
							<li><a class="rmLink" href="/resource-center/health-library">Health Library Resources</a></li>
						</ul>
					</li>
					<li><a href="/locations">Locations</a>
						<ul class="child-nav">
							<li><a class="rmLink" href="/locations/inova-spine-assessment">Inova Spine Assessment Centers</a></li>
							<li><a class="rmLink" href="/locations/inova-alexandria-hospital">Inova Alexandria Hospital</a></li>
							<li><a class="rmLink" href="/locations/inova-fairfax-hospital">Inova Fairfax Hospital</a></li>
							<li><a class="rmLink" href="/locations/inova-fair-oaks-hospital">Inova Fair Oaks Hospital</a></li>
							<li><a class="rmLink" href="/locations/inova-loudoun-hospital">Inova Loudoun Hospital</a></li>
							<li><a class="rmLink" href="/locations/inova-mount-vernon-hospital">Inova Mount Vernon Hospital</a></li>
						</ul>
					</li>
					<li><a href="/about-spine/index">About Us</a></li>
			</ul>

				</div>
                <!-- end #mobile-nav -->


                <div id="subContent">
                    <div id="subHeaderRow">
                        <div id="subHeader">
                            <img id="ctl00_BannerImage" alt="" src="/upload/images/banners/SpineInstitute/page-er-wait-times-001.jpg" />
                        </div>
                        <div id="quicklinks">
                            <h3>Quick Links</h3>
<ul>
    <li><a href="/physician-resources-spine"><img style="width: 22px; height: 22px; margin-right: 5px; margin-bottom: 4px; float: left;" alt="For Physicians icon" src="/images/design010/icon_small_physicians.png" longdesc="For Physicians icon" />For Physicians</a> </li>
    <li><a href="/?id=2992&amp;sid=22"><img style="width: 22px; height: 22px; margin-right: 5px; margin-bottom: 4px; float: left;" alt="Patient Stories icon" src="/images/design010/icon_small_patientstories.png" longdesc="Patient Stories icon" />Patient Stories</a> </li>
    <li><a href="/resource-center/online-back-in-action"><img style="width: 22px; height: 22px; margin-right: 5px; margin-bottom: 4px; float: left;" alt="Patient Education icon" src="/images/design010/icon_small_news.png" longdesc="Patient Education" />Patient Education</a> </li>
    <li><a href="https://www.inova.org/our-services/inova-neuroscience-and-spine-institute/contact-spine-program" target="_blank"><img alt="" style="width: 22px; height: 22px; margin-right: 5px; margin-bottom: 4px; float: left;" src="/images/design010/icon_small_contact.png" longdesc="Contact Us icon" />Contact Us</a> </li>
</ul>
                        </div>
                    </div>

                    <div id="contentWrap">


                        <div class="contentFull" id="contentFull">





                            <div id="printReady" class="body-copy">
                                <div id="PageContent" class="content">

                                    <style type="text/css">
    .wait-time-banner {
    text-align: center;
    }
    .wait-time-banner img {
    width: 100%;
    height: auto;
    }
    /*------ HIDES VIRTUAL CARE TILE -------*/
    .col-md-3.col-xs-12.border-left-bottom.er-wait-cell.ondemand {
    display:  none!important;
    }
</style>
<div class="wait-time-banner"><img alt="" src="/upload/images/Banners/er_wait_banner.jpg" /></div>
<h1><span style="font-size: 24px;">Average* wait times for <a href="/healthcare-services/emergency-services/locations/index.jsp">Inova emergency rooms</a></span></h1>
<!--
<div style="font-size: 12px; float: right; margin-left: 20px; width: 160px;"><strong>Average* Wait Times for Inova Emergency Rooms<br />
wait times are</strong><br />
<span style="font-weight: bold; font-size: 16px; color: #e39303;">updated every<br />
10 minutes.</span></div>
<p style="font-weight: bold; font-size: 11px; color: red;">If you are experiencing a medical emergency, please call 911 immediately.</p>
<p style="font-size: 11px;">For other times when you need to visit one of our emergency rooms, we have provided estimates of our current wait times. These times are provided for informational purposes only and cannot be guaranteed upon arrival. They may be subject to change without notice.</p>
--><img alt="" src="http://ds.reson8.com/insights.gif?rand=[cache_buster]&amp;t=0&amp;pixt=resonate&amp;advkey=0013000001AeNwJAAV&amp;opptykey=IHES0915A&amp;evkey=131095&amp;evtype=custom" width="1" height="1" style="border-width: 0px; border-style: solid;" />
<!--––VOICE MEDIA PIXELS 9/11/18 ––-->
<img alt="" src="https://secure.adnxs.com/px?id=1030458&amp;seg=14557005&amp;t=2" width="1" height="1" />
<img alt="" src="https://ad.doubleclick.net/ddm/activity/src=8117457;type=invmedia;cat=xydjan21;dc_lat=;dc_rdid=;tag_for_child_directed_treatment=;tfua=;npa=;ord=1?" width="1" height="1" />



    <div class="container">
        <div class="row">
            <div class="col-md-3 col-xs-12 border-left-bottom er-wait-cell">
			<div class='er-location'><a href='https://www.google.com/maps/place/Inova+Alexandria+Hospital/@38.834943,-77.222034,11.75z/data=!4m5!1m2!2m1!1sinova+alexandria+hospital!3m1!1s0x89b7b3d831179359:0xc1e644d8b3c177e' target='_blank'><h3>43<abbr title='minute' class='min-abb'>min</abbr></h3><div class='location-details'><span class='location-name'>Inova Alexandria Hospital</span><div class='location-address'><span class='location-icon'></span><span class='location-address-line1'>4320 Seminary Rd</span><span class='location-address-line2'>Alexandria, VA 22304</span></div></div><div class='clearer'></div></a></div>
            </div>
            <div class="col-md-3 col-xs-12 border-left-bottom er-wait-cell">
            <div class='er-location'><a href='https://www.google.com/maps/place/Inova+Emergency+Care+Center+-+HealthPlex+Ashburn/@38.99472,-77.4829217,17z/data=!3m1!4b1!4m2!3m1!1s0x89b63f486f0fc1af:0x61d4ebe95024b708' target='_blank'><h3>18<abbr title='minute' class='min-abb'>min</abbr></h3><div class='location-details'><span class='location-name'>Inova HealthPlex Emergency Room &mdash; Ashburn</span><span class='small-text'>A service of<br />Inova Loudoun Hospital</span><div class='location-address'><span class='location-icon'></span><span class='location-address-line1'>22505 Landmark Court</span><span class='location-address-line2'>Ashburn, VA 20148</span></div></div><div class='clearer'></div></a></div>
            </div>
            <div class="col-md-3 col-xs-12 border-left-bottom er-wait-cell">
            <div class='er-location'><a href='https://www.google.com/maps/place/Inova+Emergency+Room+-+HealthPlex+Franconia%2FSpringfield/@38.7674,-77.1629207,17z/data=!3m1!4b1!4m2!3m1!1s0x89b7ad65377d09f3:0x8ac833dce0c6bebf' target='_blank'><h3>10<abbr title='minute' class='min-abb'>min</abbr></h3><div class='location-details'><span class='location-name'>Inova HealthPlex Emergency Room &mdash; Franconia/Springfield</span><span class='small-text'>A service of<br />Inova Alexandria Hospital</span><div class='location-address'><span class='location-icon'></span><span class='location-address-line1'>6355 Walker Ln</span><span class='location-address-line2'>Alexandria, VA 22310</span></div></div><div class='clearer'></div></a></div>
            </div>
            <div class="col-md-3 col-xs-12 border-bottom er-wait-cell">
            <div class='er-location'><a href='https://www.google.com/maps/place/Inova+Fairfax+Hospital/@38.8574974,-77.2306172,17z/data=!3m1!4b1!4m2!3m1!1s0x89b64c9ac6b9265f:0xaf05ab4d8e208180' target='_blank'><h3>15<abbr title='minute' class='min-abb'>min</abbr></h3><div class='location-details'><span class='location-name'>Inova Fairfax Hospital</span><div class='location-address'><span class='location-icon'></span><span class='location-address-line1'>3300 Gallows Rd</span><span class='location-address-line2'>Falls Church, VA 22042</span></div></div><div class='clearer'></div></a></div>
		</div>
        </div>
        <div class="row">
            <div class="col-md-3 col-xs-12 border-left-bottom er-wait-cell">
            <div class='er-location'><a href='https://www.google.com/maps/place/Inova+Children%27s+Hospital/@38.8575057,-77.2306172,17z/data=!4m5!1m2!2m1!1sInova+Children%E2%80%99s+Emergency+Room+!3m1!1s0x89b64c9ac6b9265f:0x9f5b1b61e3b42608' target='_blank'><h3>19<abbr title='minute' class='min-abb'>min</abbr></h3><div class='location-details'><span class='location-name'>Inova Children's Emergency Room &mdash; Inova Fairfax Medical Campus</span><div class='location-address'><span class='location-icon'></span><span class='location-address-line1'>3300 Gallows Rd</span><span class='location-address-line2'>Falls Church, VA 22042</span></div></div><div class='clearer'></div></a></div>
            </div>
            <div class="col-md-3 col-xs-12 border-left-bottom er-wait-cell">
            <div class='er-location'><a href='https://www.google.com/maps/place/Inova+Emergency+Room+-+Fairfax/@38.8521789,-77.2890718,14z/data=!4m5!1m2!2m1!1sInova+Emergency+Care+Center+%E2%80%94+Fairfax!3m1!1s0x89b64ef4b9156135:0x5a597bbf52a6aac6' target='_blank'><h3>0<abbr title='minute' class='min-abb'>min</abbr></h3><div class='location-details'><span class='location-name'>Inova Emergency Room &mdash; Fairfax</span><span class='small-text'>A service of<br />Inova Fairfax Hospital</span><div class='location-address'><span class='location-icon'></span><span class='location-address-line1'>4315 Chain Bridge Rd</span><span class='location-address-line2'>Fairfax, VA 22030</span></div></div><div class='clearer'></div></a></div>
            </div>
            <div class="col-md-3 col-xs-12 border-left-bottom er-wait-cell">
            <div class='er-location'><a href='https://www.google.com/maps/place/Inova+Emergency+Room+-+Reston%2FHerndon/@38.965411,-77.3588877,17z/data=!3m1!4b1!4m2!3m1!1s0x0:0x50769888f7d2cc7b?hl=en' target='_blank'><h3>21<abbr title='minute' class='min-abb'>min</abbr></h3><div class='location-details'><span class='location-name'>Inova Emergency Room &mdash; Reston/Herndon</span><span class='small-text'>A service of<br />Inova Fairfax Hospital</span><div class='location-address'><span class='location-icon'></span><span class='location-address-line1'>11901 Baron Cameron Ave</span><span class='location-address-line2'>Reston, VA 20190</span></div></div><div class='clearer'></div></a></div>
            </div>
            <div class="col-md-3 col-xs-12 border-bottom er-wait-cell">
            <div class='er-location'><a href='https://www.google.com/maps/place/Inova+Fair+Oaks+Hospital+Emergency+Room/@38.8849422,-77.3845307,17z/data=!3m1!4b1!4m2!3m1!1s0x89b64600f75a0547:0x59200abd2905059c' target='_blank'><h3>12<abbr title='minute' class='min-abb'>min</abbr></h3><div class='location-details'><span class='location-name'>Inova Fair Oaks Hospital</span><div class='location-address'><span class='location-icon'></span><span class='location-address-line1'>3600 Joseph Siewick Dr</span><span class='location-address-line2'>Fairfax, VA 22033</span></div></div><div class='clearer'></div></a></div>
		</div>
        </div>
        <div class="row">
            <div class="col-md-3 col-xs-12 border-left-bottom er-wait-cell">
            <div class='er-location'><a href='https://www.google.com/maps/place/Inova+Loudoun+Hospital/@39.075,-77.5455828,12z/data=!4m5!1m2!2m1!1sInova+Loudoun+Hospital!3m1!1s0x89b63c0c38ea7219:0x2bedd03869965bae' target='_blank'><h3>2<abbr title='minute' class='min-abb'>min</abbr></h3><div class='location-details'><span class='location-name'>Inova Loudoun Hospital</span><div class='location-address'><span class='location-icon'></span><span class='location-address-line1'>44045 Riverside Pkwy</span><span class='location-address-line2'>Leesburg, VA 20176</span></div></div><div class='clearer'></div></a></div>
            </div>
            <div class="col-md-3 col-xs-12 border-left-bottom er-wait-cell">
            <div class='er-location'><a href='https://www.google.com/maps/place/Inova+Loudoun+Hospital+Children%E2%80%99s+Emergency+Room/@39.0740491,-77.4797237,17z/data=!4m5!1m2!2m1!1sInova+Loudoun+Hospital+Children%E2%80%99s+Emergency+Room+!3m1!1s0x89b63c0c3dca5fbd:0x22b2ee6d5332096c' target='_blank'><h3>13<abbr title='minute' class='min-abb'>min</abbr></h3><div class='location-details'><span class='location-name'>Inova Loudoun Hospital Children's Emergency Room</span><div class='location-address'><span class='location-icon'></span><span class='location-address-line1'>44045 Riverside Parkway</span><span class='location-address-line2'>Leesburg, VA  20176</span></div></div><div class='clearer'></div></a></div>
            </div>
            <div class="col-md-3 col-xs-12 border-left-bottom er-wait-cell">
            <div class='er-location'><a href='https://www.google.com/maps/place/Inova+Emergency+Room+-+Leesburg/@39.1197986,-77.5710414,17z/data=!4m5!1m2!2m1!1sInova+Emergency+Care+Center+%E2%80%94+Leesburg!3m1!1s0x89b617da16c72ae5:0x644520b96b44c963' target='_blank'><h3>0<abbr title='minute' class='min-abb'>min</abbr></h3><div class='location-details'><span class='location-name'>Inova Emergency Room &mdash; Leesburg</span><div class='location-address'><span class='location-icon'></span><span class='location-address-line1'>224 Cornwall St, NW</span><span class='location-address-line2'>Leesburg, VA 20176</span></div></div><div class='clearer'></div></a></div>
            </div>
            <div class="col-md-3 col-xs-12 border-bottom er-wait-cell">
            <div class='er-location'><a href='https://www.google.com/maps/place/Inova+Mount+Vernon+Hospital/@38.7397212,-77.0801607,17z/data=!4m5!1m2!2m1!1sInova+Mount+Vernon+Hospital!3m1!1s0x89b7ae42347d4d6b:0xf3f535b0e643e311' target='_blank'><h3>3<abbr title='minute' class='min-abb'>min</abbr></h3><div class='location-details'><span class='location-name'>Inova Mount Vernon Hospital</span><div class='location-address'><span class='location-icon'></span><span class='location-address-line1'>2501 Parker's Ln</span><span class='location-address-line2'>Alexandria, VA 22306</span></div></div><div class='clearer'></div></a></div>
            </div>
        </div>
        <div class="row">
            <div class="col-md-3 col-xs-12 border-left-bottom er-wait-cell">
                <div class='er-location'><a href='https://www.google.com/maps/place/Inova+Emergency+Room+-+HealthPlex+Lorton/@38.7026102,-77.229147,17z/data=!4m5!1m2!2m1!1sInova+Emergency+Care+Center+%E2%80%94Lorton!3m1!1s0x89b65362ed66d82f:0xe9b1f94d47732549' target='_blank'><h3>2<abbr title='minute' class='min-abb'>min</abbr></h3><div class='location-details'><span class='location-name'>Inova Emergency Room &mdash; Lorton</span><span class='small-text'>A service of<br />Inova Mount Vernon Hospital</span><div class='location-address'><span class='location-icon'></span><span class='location-address-line1'>9321 Sanger Street</span><span class='location-address-line2'>Lorton, VA  22079</span></div></div><div class='clearer'></div></a></div></div>

            <div class="col-md-3 col-xs-12 border-left-bottom er-wait-cell ondemand">
                <div class="ondemand-promo-box">


                    <div class="clearer"></div>
                </div>
            </div>
            <div class="col-md-3 col-xs-12 border-left-bottom er-wait-cell ucc">
                <div class="ondemand-promo-box">


                    <div class="clearer"></div>
                </div>

            <div class="clearer"></div>
        </div>
    </div>



        <div class="er-wait-list-row">
			<!-- STATIC -->
			<div class="er-location disclaimer">
            <p>
                *<strong>How are Emergency Room wait times calculated?</strong> Emergency Room wait times represent the average estimated length of
				time over a one hour period from registration to "assignment of provider" (a doctor, nurse practitioner or physician assistant).
				The wait times displayed on this site are refreshed at least every 30 minutes. Many circumstances can affect wait
				times – for example, patients arriving by ambulance or with life-threatening injuries or illnesses. Patients with
				life-threatening conditions will be seen before those with less-serious problems or ailments. These times are
				provided for informational purposes only and cannot be guaranteed upon arrival. They may be subject to change
				without notice. When NA is displayed, the wait time is not available.
            </p>
			</div>
			<div class="clearer"></div>
		</div>

		<div id="top-footer-relocate" class="mobile-menu page-related-menu expanded"></div>
		<div id="share-relocate"></div>
		<div class="clearer"></div>


                                </div>
                            </div>



                        </div>


                        <!--ContentLayoutColumn-->

                    </div>
                    <!--Contentwrap-->
                </div>
                <!--<div id="subContent"> open in sub top-->




                <!--</div> -->
            </div>

 <div id="row4">
    <div id="footer010">
        <div class="footerCol">
            <h3>Quick Links</h3>
<ul>
    <li><a href="/?id=2980&amp;sid=22">Meet the Team</a> </li>
    <li><a href="/?id=3030&amp;sid=22">Spine Conditions</a></li>
    <li><a href="/?id=2982&amp;sid=22">Treatment Options</a> </li>
    <li><a href="/?id=2983&amp;sid=22">Resource Center</a> </li>
    <li><a href="/?id=2984&amp;sid=22">Locations</a> </li>
    <li><a href="/?id=2985&amp;sid=22">About Us</a> </li>
    <li><a href="http://www.inovaspine.org/about-spine/give-a-gift-to-inova-spine">Donate</a></li>
    <li><a href="https://www.inova.org/about-inova/subscribe" target="_blank">Subscribe to Newsletter</a></li>
</ul>
<p>​</p>
<p>​</p>
        </div>
        <div class="footerCol">
            <h3>Patient Information</h3>
<ul>
    <li><a href="http://www.inova.org/billing">Billing</a></li>
    <li><a href="http://www.inova.org/patient-and-visitor-information/financial-help/index.jsp">Financial Help</a></li>
    <li><a href="http://www.inova.org/patient-and-visitor-information/making-healthcare--decisions/index.jsp">Making Healthcare Decisions</a></li>
    <li><a href="http://www.inova.org/patient-and-visitor-information/privacy-and-compliance/index.jsp">Privacy and Compliance</a></li>
    <li><a href="http://www.inova.org/patient-visitor/nondiscrimination-policy">Nondiscrimination Policy</a></li>
    <li><a href="http://www.inova.org/patient-and-visitor-information/medical-records/index.jsp">Request Medical  Records</a></li>
</ul>
        </div>
        <div class="footerCol">
            <h3>
Locations
</h3>
<ul>
    <li>
    <a href="/locations/inova-alexandria-hospital">Inova Alexandria Hospital</a>
    </li>
    <li>
    <a href="/locations/inova-fairfax-hospital">Inova Fairfax Hospital</a>
    </li>
    <li>
    <a href="/locations/inova-fair-oaks-hospital">Inova Fair Oaks Hospital</a>
    </li>
    <li>
    <a href="/locations/inova-loudoun-hospital">Inova Loudoun Hospital</a>
    </li>
    <li>
    <a href="/locations/inova-mount-vernon-hospital">Inova Mount Vernon Hospital</a>
    </li>
    <li>
    <a href="/locations/inova-spine-assessment">Spine Assessment Centers at Inova Urgent Care Centers</a></li>
</ul>
        </div>
        <div class="footerCol footerLast">
            <div id="social"><a href="http://twitter.com/InovaHealth" class="twitter" target="_blank"><img alt="" width="31" height="31" src="/images/design010/iconTwitter.jpg " /></a> <a href="http://www.youtube.com/InovaHealthSystem" class="youtube" target="_blank"><img alt="" width="31" height="31" src="/images/design010/iconYouTube.jpg" /></a> <a href="http://www.facebook.com/InovaHealth" class="facebook" target="_blank"><img alt="" width="31" height="31" src="/images/design010/iconFacebook.jpg" /></a> <a href="http://www.carepages.com/inova" class="carepages" target="_blank"><img alt="" width="31" height="31" src="/images/design010/iconCarePages.jpg" /></a> </div>
<ul>
    <li><a href="https://www.inova.org/our-services/inova-neuroscience-and-spine-institute/contact-spine-program" target="_blank">Contact Us</a></li>
    <li><a href="/sitemap/?sid=22">Sitemap</a> </li>
    <li><a href="/?id=1&amp;sid=1">Inova.org</a></li>
</ul>
        </div>
    </div>
    <br />
    <div class="CopyrightMiddle">
        <a href="/admin/" style="text-decoration: none; color: #333;" onclick="window.open('/admin/'); return false;">&copy;</a>
        2022 Inova Spine
        <a href="http://www.inovachildrens.org/admin/pages/?PageID=1014" id="ctl00_ContentFooter_pencilLink" target="_blank">
            <img alt="Pencil" src="/images/pencil.gif" width="8" height="8" style="border: none;" />
        </a>
    </div>
</div>


	<script type="text/javascript">setTimeout(function(){var a=document.createElement("script");var b=document.getElementsByTagName("script")[0];a.src=document.location.protocol+"//dnn506yrbagrg.cloudfront.net/pages/scripts/0013/0969.js?"+Math.floor(new Date().getTime()/3600000);a.async=true;a.type="text/javascript";b.parentNode.insertBefore(a,b)}, 1);</script>
 <script src="/scripts/design010/dist/jquery.slicknav.js"></script>
        <script type="text/javascript">
        $(document).ready(function(){
            $('#mobile-menu').slicknav();

            $(".slicknav_nav .slicknav_row").on('click',function(){
                $(this).parent().children('ul.slicknav_hidden').slideToggle();
                //console.log($(this).parent());
            });

            $("div.slicknav_menu").prepend("<div data-toggle='modal' data-target='#globalHeaderMenuModal' style='float:right!important;color:white;' class='imc-searchoverlay-app slicknav_btn'>SEARCH<i style='font-size:12px;line-height:16px;margin-left:5px;margin-right:5px;' class='fa fa-fw fa-search'></i></div>");

        });
        </script>


<script type="text/javascript">
//<![CDATA[
Sys.Application.add_init(function() {
    $create(Telerik.Web.UI.RadMenu, {"_childListElementCssClass":null,"_skin":"Design010","clientStateFieldID":"ctl00_PrimaryNavigation_ClientState","collapseAnimation":"{\"duration\":450}","enableAutoScroll":true,"expandAnimation":"{\"duration\":450}","itemData":[{"items":[{"navigateUrl":"/our-spine-team/medical-directors/"},{"navigateUrl":"/our-spine-team/physicians"},{"navigateUrl":"/our-spine-team/care-coordination-team"}],"navigateUrl":"/our-spine-team"},{"isSeparator":true},{"items":[{"navigateUrl":"/spine-conditions/types-of-conditions"},{"navigateUrl":"/spine-conditions/symptoms"},{"navigateUrl":"/spine-conditions/diagnosing"},{"navigateUrl":"/spine-conditions-resources/"}],"navigateUrl":"/spine-conditions"},{"isSeparator":true},{"items":[{"items":[{"navigateUrl":"/treatment-options/surgery/minimally-invasive-spine-surgery"}],"navigateUrl":"/treatment-options/surgery"},{"navigateUrl":"/treatment-options/physical-therapy"},{"navigateUrl":"/treatment-options/pain-management"}],"navigateUrl":"/treatment-options/"},{"isSeparator":true},{"items":[{"navigateUrl":"/resource-center/blog-posts"},{"navigateUrl":"/patient-resources/preparing-for-surgery-or-treatment"},{"navigateUrl":"/spine-doctor-videos"},{"navigateUrl":"/resource-center/forms"},{"navigateUrl":"/resource-center/online-back-in-action"},{"navigateUrl":"/resource-center/health-library"}],"navigateUrl":"/resource-center/prep-for-surgery"},{"isSeparator":true},{"navigateUrl":"/virtual-visits"},{"isSeparator":true},{"items":[{"navigateUrl":"/locations/inova-spine-assessment"},{"navigateUrl":"/locations/inova-alexandria-hospital"},{"navigateUrl":"/locations/inova-fairfax-hospital"},{"navigateUrl":"/locations/inova-fair-oaks-hospital"},{"navigateUrl":"/locations/inova-loudoun-hospital"},{"navigateUrl":"/locations/inova-mount-vernon-hospital"}],"navigateUrl":"/locations"},{"isSeparator":true},{"navigateUrl":"/about-spine/index"}]}, null, null, $get("ctl00_PrimaryNavigation"));
});
//]]>
</script>
</form>


	<!--Centretek search script-->
	<script type="text/javascript" src="https://assets.inova-search-drupal.com/centretek_search.js"></script>
	<script type="text/javascript" src="/Scripts/inovaspine.js"></script>
	<!--End Centretek search script-->
    </body>
</html>
`)

func TestParseWaits(t *testing.T) {
	waits, err := ParseWaits(htmlSample)
	if err != nil {
		t.Fatalf("Unexpected parse error: %v", err)
	}

	if expect, got := 13, len(waits); expect != got {
		t.Fatalf("Expected %d waits; got %d", expect, got)
	}
}
