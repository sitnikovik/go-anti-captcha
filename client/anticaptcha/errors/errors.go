package errors

import "fmt"

const (
	keyDoesNotExistID  = 1
	keyDoesNotExistTxt = "account authorization key not found in the system"

	noSlotAvailableID  = 2
	noSlotAvailableTxt = "no idle captcha workers are available at the moment, please try a bit later or try increasing your bid"

	zeroCaptchaFilesizeID  = 3
	zeroCaptchaFilesizeTxt = "the size of the captcha file is too small"

	tooBigCaptchaFilesizeID  = 4
	tooBigCaptchaFilesizeTxt = "the size of the captcha file is too big"

	zeroBalanceID  = 10
	zeroBalanceTxt = "account has zero or negative balance"

	ipNotAllowedID  = 11
	ipNotAllowedTxt = "request sent from the IP that is not allowed with current account key"

	captchaUnsolvableID  = 12
	captchaUnsolvableTxt = "captcha could not be solved"

	noSuchMethodID  = 14
	noSuchMethodTxt = "request made to API with a method that does not exist"

	imageTypeNotSupportedID  = 15
	imageTypeNotSupportedTxt = "image type not supported"

	noSuchCaptchaID  = 16
	noSuchCaptchaTxt = "task not found or not accessible"

	ipBlockedID  = 21
	ipBlockedTxt = "your IP is blocked due to improper use of API"

	taskAbsentID  = 22
	taskAbsentTxt = "task property is empty or not set in the createTask method"

	taskNotSupportedID  = 23
	taskNotSupportedTxt = "task type is not supported or typed incorrectly"

	incorrectSessionDataID  = 24
	incorrectSessionDataTxt = "some of the required values for successive user emulation are missing"

	proxyConnectRefusedID  = 25
	proxyConnectRefusedTxt = "could not connect to task proxy, connection refused"

	proxyConnectTimeoutID  = 26
	proxyConnectTimeoutTxt = "could not connect to task proxy, connection timed out"

	proxyReadTimeoutID  = 27
	proxyReadTimeoutTxt = "reading timeout of task's proxy"

	proxyBannedID  = 28
	proxyBannedTxt = "proxy IP banned by target service"

	proxyTransparentID  = 29
	proxyTransparentTxt = "proxy must be non-transparent to hide our server IP"

	recaptchaTimeoutID  = 30
	recaptchaTimeoutTxt = "recaptcha task timeout, probably due to slow proxy server or Google server"

	recaptchaInvalidSitekeyID  = 31
	recaptchaInvalidSitekeyTxt = "captcha provider reported that the site key is invalid"

	recaptchaInvalidDomainID  = 32
	recaptchaInvalidDomainTxt = "captcha provider reported that the domain for this site key is invalid"

	recaptchaOldBrowserID  = 33
	recaptchaOldBrowserTxt = "captcha provider reported that the browser user-agent is not compatible with their javascript"

	tokenExpiredID  = 34
	tokenExpiredTxt = "captcha provider server reported that the additional variable token has expired"

	proxyHasNoImageSupportID  = 35
	proxyHasNoImageSupportTxt = "proxy does not support transfer of image data from Google servers"

	proxyIncompatibleHttpVersionID  = 36
	proxyIncompatibleHttpVersionTxt = "proxy does not support long GET requests with length about 2000 bytes and does not support SSL connections"

	proxyNotAuthorisedID  = 49
	proxyNotAuthorisedTxt = "proxy login and password are incorrect"

	proxyInvalidKeyTypeID  = 51
	proxyInvalidKeyTypeTxt = "passed sitekey is from another Recaptcha type. Try solve it as V2, V2-invisible or V3."

	failedLoadingWidgetID  = 52
	failedLoadingWidgetTxt = "Could not load captcha provider widget in worker browser. Please try sending a new task."

	visibleRecaptchaID  = 53
	visibleRecaptchaTxt = "Attempted solution of usual Recaptcha V2 as Recaptcha V2 invisible. Remove flag 'isInvisible' from the API payload."

	allWorkersFilteredID  = 54
	allWorkersFilteredTxt = "no workers left that were not filtered by reportIncorrectRecaptcha method."

	accountSuspendedID  = 55
	accountSuspendedTxt = "the system has suspended your account for a significant reason. Contact support for details."

	templateNotFoundID  = 56
	templateNotFoundTxt = "AntiGate template not found by its name during task creation"

	taskCanceledID  = 57
	taskCanceledTxt = "AntiGate task was canceled by worker"
)

// ByErrorID makes error by id.
// See more https://anti-captcha.com/apidoc/errors
func ByErrorID(id int) error {
	m := map[int]string{
		keyDoesNotExistID:              keyDoesNotExistTxt,
		noSlotAvailableID:              noSlotAvailableTxt,
		zeroCaptchaFilesizeID:          zeroCaptchaFilesizeTxt,
		tooBigCaptchaFilesizeID:        tooBigCaptchaFilesizeTxt,
		zeroBalanceID:                  zeroBalanceTxt,
		ipNotAllowedID:                 ipNotAllowedTxt,
		captchaUnsolvableID:            captchaUnsolvableTxt,
		noSuchMethodID:                 noSuchMethodTxt,
		imageTypeNotSupportedID:        imageTypeNotSupportedTxt,
		noSuchCaptchaID:                noSuchCaptchaTxt,
		ipBlockedID:                    ipBlockedTxt,
		taskAbsentID:                   taskAbsentTxt,
		taskNotSupportedID:             taskNotSupportedTxt,
		incorrectSessionDataID:         incorrectSessionDataTxt,
		proxyConnectRefusedID:          proxyConnectRefusedTxt,
		proxyConnectTimeoutID:          proxyConnectTimeoutTxt,
		proxyReadTimeoutID:             proxyReadTimeoutTxt,
		proxyBannedID:                  proxyBannedTxt,
		proxyTransparentID:             proxyTransparentTxt,
		recaptchaTimeoutID:             recaptchaTimeoutTxt,
		recaptchaInvalidSitekeyID:      recaptchaInvalidSitekeyTxt,
		recaptchaInvalidDomainID:       recaptchaInvalidDomainTxt,
		recaptchaOldBrowserID:          recaptchaOldBrowserTxt,
		tokenExpiredID:                 tokenExpiredTxt,
		proxyHasNoImageSupportID:       proxyHasNoImageSupportTxt,
		proxyIncompatibleHttpVersionID: proxyIncompatibleHttpVersionTxt,
		proxyNotAuthorisedID:           proxyNotAuthorisedTxt,
		proxyInvalidKeyTypeID:          proxyInvalidKeyTypeTxt,
		failedLoadingWidgetID:          failedLoadingWidgetTxt,
		visibleRecaptchaID:             visibleRecaptchaTxt,
		allWorkersFilteredID:           allWorkersFilteredTxt,
		accountSuspendedID:             accountSuspendedTxt,
		templateNotFoundID:             templateNotFoundTxt,
		taskCanceledID:                 taskCanceledTxt,
	}

	txt, ok := m[id]
	if !ok {
		txt = "unknown error"
	}

	return fmt.Errorf("code %d: %s. See more https://anti-captcha.com/apidoc/errors", id, txt)
}

// IsUnauthorized returns true if error is about unauthorized key
func IsUnauthorized(err error) bool {
	return err.Error() == ByErrorID(keyDoesNotExistID).Error()
}

// IsTaskNotFound returns true if error is about missing key
func IsTaskNotFound(err error) bool {
	return err.Error() == ByErrorID(noSuchCaptchaID).Error()
}

// IsImageTypeNotSupported returns true if error is about image type not supported
func IsImageTypeNotSupported(err error) bool {
	return err.Error() == ByErrorID(imageTypeNotSupportedID).Error()
}

// IsNoSlotAvailable returns true if error is about no slot available
func IsNoSlotAvailable(err error) bool {
	return err.Error() == ByErrorID(noSlotAvailableID).Error()
}

// IsZeroBalance returns true if error is about zero balance
func IsZeroBalance(err error) bool {
	return err.Error() == ByErrorID(zeroBalanceID).Error()
}

// IsIpNotAllowed returns true if error is about IP not allowed
func IsIpNotAllowed(err error) bool {
	return err.Error() == ByErrorID(ipNotAllowedID).Error()
}

// IsCaptchaUnsolvable returns true if error is about captcha unsolvable
func IsCaptchaUnsolvable(err error) bool {
	return err.Error() == ByErrorID(captchaUnsolvableID).Error()
}
