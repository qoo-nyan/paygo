package paygo

var (
	EndpointPayPayApp string = "https://app4.paypay.ne.jp/"

	EndpointAPIV1 string = EndpointPayPayApp + "bff/v1/"
	EndpointAPIV2 string = EndpointPayPayApp + "bff/v2"
	EndpointAPIV3 string = EndpointPayPayApp + "bff/v3/"

	EndpotintBalanceInfo                                  = EndpointAPIV1 + "getBalanceInfo"
	EndpotintWalletWidgetInfo                             = EndpointAPIV1 + "getWalletWidgetInfo"
	EndpotintPayLaterCcAvailableAmount                    = EndpointAPIV1 + "getPayLaterCcAvailableAmount"
	EndpotintPointBalance                                 = EndpointAPIV1 + "getPointBalance"
	EndpotintUserProfile                                  = EndpointAPIV1 + "getUserProfile"
	EndpotintChangeUserProfile                            = EndpointAPIV1 + "changeUserProfile"
	EndpotintStatementSummary                             = EndpointAPIV1 + "getStatementSummary"
	EndpotintSetUserDefinedLimitType                      = EndpointAPIV1 + "setUserDefinedLimitType"
	EndpotintDeleteTemporaryMailAddress                   = EndpointAPIV1 + "deleteTemporaryMailAddress"
	EndpotintReAuthDisplayInfo                            = EndpointAPIV1 + "getReAuthDisplayInfo"
	EndpotintValidatePayPayId                             = EndpointAPIV1 + "validatePayPayId"
	EndpotintDeleteP2PUserSearchHistory                   = EndpointAPIV1 + "deleteP2PUserSearchHistory"
	EndpotintV1PayoutDisplayInfo                          = EndpointAPIV1 + "getPayoutDisplayInfo"
	EndpotintYahooWalletToken                             = EndpointAPIV1 + "getYahooWalletToken"
	EndpotintP2PMessageList                               = EndpointAPIV1 + "getYahooWalletToken"
	EndpotintContinuousPaymentDisplayInfo                 = EndpointAPIV1 + "getContinuousPaymentDisplayInfo"
	EndpotintExecuteReAuthPayment                         = EndpointAPIV1 + "executeReAuthPayment"
	EndpotintChannelDisplayInfo                           = EndpointAPIV1 + "getChannelDisplayInfo"
	EndpotintAgreeSimilarTransaction                      = EndpointAPIV1 + "agreeSimilarTransaction"
	EndpotintCancelP2PRequestMoney                        = EndpointAPIV1 + "cancelP2PRequestMoney"
	EndpotintNearbyStoreList                              = EndpointAPIV1 + "getNearbyStoreList"
	EndpotintShareFeed                                    = EndpointAPIV1 + "shareFeed"
	EndpotintCreateP2PUserSearchHistory                   = EndpointAPIV1 + "createP2PUserSearchHistory"
	EndpotintPaymentDetailDisplayInfo                     = EndpointAPIV1 + "getPaymentDetailDisplayInfo"
	EndpotintRemoveUserDefinedLimit                       = EndpointAPIV1 + "removeUserDefinedLimit"
	EndpotintAuthorizePassword                            = EndpointAPIV1 + "authorizePassword"
	EndpotintOrder                                        = EndpointAPIV1 + "getOrder"
	EndpotintOrderByOrderId                               = EndpointAPIV1 + "getOrderByOrderId"
	EndpotintP2PThemeList                                 = EndpointAPIV1 + "getP2PThemeList"
	EndpotintUnLikeFeed                                   = EndpointAPIV1 + "unlikeFeed"
	EndpotintCreateP2PUserChannel                         = EndpointAPIV1 + "createP2PUserChannel"
	EndpotintAcceptP2PRequestMoney                        = EndpointAPIV1 + "acceptP2PRequestMoney"
	EndpotintUpdatePrioritizedPaymentMethodsConfiguration = EndpointAPIV1 + "updatePrioritizedPaymentMethodsConfiguration"
	EndpotintSignOutAndDeleteAllDevices                   = EndpointAPIV1 + "signOutAndDeleteAllDevices"
	EndpotintLikeFeed                                     = EndpointAPIV1 + "likeFeed"
	EndpotintDeletePayoutMethod                           = EndpointAPIV1 + "deletePayoutMethod"
	EndpotintFavoriteSmartFunctionList                    = EndpointAPIV1 + "getFavoriteSmartFunctionList"
	EndpotintExecutePayout                                = EndpointAPIV1 + "executePayout"
	EndpotintPay2BalanceCancelledPendingCashBackHistory   = EndpointAPIV1 + "getPay2BalanceCancelledPendingCashBackHistory"
	EndpotintSignOutAndDeleteDevice                       = EndpointAPIV1 + "signOutAndDeleteDevice"
	EndpotintServiceStatus                                = EndpointAPIV1 + "getServiceStatus"
	EndpotintUpdateDeviceLockStatus                       = EndpointAPIV1 + "updateDeviceLockStatus"
	EndpotintCreateP2PSession                             = EndpointAPIV1 + "createP2PSession"
	EndpotintSendForgetLink                               = EndpointAPIV1 + "sendForgetLink"
	EndpotintYconnectLogin                                = EndpointAPIV1 + "yconnectLogin"
	EndpotintFeedList                                     = EndpointAPIV1 + "getFeedList"
	EndpotintP2POrder                                     = EndpointAPIV1 + "getP2POrder"
	EndpotintP2PUserSearchHistory                         = EndpointAPIV1 + "getP2PUserSearchHistory"
	EndpotintCreatePayPaySignInCode                       = EndpointAPIV1 + "createPayPaySignInCode"
	EndpotintUpdateGvAutoSelectConfiguration              = EndpointAPIV1 + "updateGvAutoSelectConfiguration"
	EndpotintRejectP2PSendMoney                           = EndpointAPIV1 + "rejectP2PSendMoney"
	EndpotintAuthorizeYconnect                            = EndpointAPIV1 + "authorizeYconnect"
	EndpotintAgreeToPrivacyPolicy                         = EndpointAPIV1 + "agreeToPrivacyPolicy"
	EndpotintUserDefinedLimits                            = EndpointAPIV1 + "getUserDefinedLimits"
	EndpotintTopupOrder                                   = EndpointAPIV1 + "getTopupOrder"
	EndpotintGiftCardOrder                                = EndpointAPIV1 + "getGiftCardOrder"
	EndpotintRegisterPayPayId                             = EndpointAPIV1 + "registerPayPayId"
	EndpotintChangePaymentMethodState                     = EndpointAPIV1 + "changePaymentMethodState"
	EndpotintCalculateMynaSendApplyInfo                   = EndpointAPIV1 + "calculateMynaSendApplyInfo"
	EndpotintChangeDateOfBirth                            = EndpointAPIV1 + "changeDateOfBirth"
	EndpotintCreateP2PCode                                = EndpointAPIV1 + "createP2PCode"
	EndpotintChangeAvatarImage                            = EndpointAPIV1 + "changeAvatarImage"
	EndpotintPendingContinuousPaymentList                 = EndpointAPIV1 + "getPendingContinuousPaymentList"
	EndpotintP2PRequestOrder                              = EndpointAPIV1 + "getP2PRequestOrder"
	EndpotintTpointToken                                  = EndpointAPIV1 + "getTpointToken"
	EndpotintContinuousPaymentHistory                     = EndpointAPIV1 + "getContinuousPaymentHistory"
	EndpotintRegisterGiftCard                             = EndpointAPIV1 + "registerGiftCard"
	EndpotintPaymentHistoryFilterCondition                = EndpointAPIV1 + "getPaymentHistoryFilterCondition"
	EndpotintServiceLinkageDisplayInfo                    = EndpointAPIV1 + "getServiceLinkageDisplayInfo"
	EndpotintExecutePreAuthPayment                        = EndpointAPIV1 + "executePreAuthPayment"
	EndpotintLocalizedString                              = EndpointAPIV1 + "getLocalizedString"
	EndpotintChangeSearchablePayPayId                     = EndpointAPIV1 + "changeSearchablePayPayId"
	EndpotintUpdateOpenPaymentClientStatus                = EndpointAPIV1 + "updateOpenPaymentClientStatus"
	EndpotintSmartLoginSession                            = EndpointAPIV1 + "getSmartLoginSession"
	EndpotintExecutePayment                               = EndpointAPIV1 + "executePayment"
	EndpotintFollowChannel                                = EndpointAPIV1 + "followChannel"
	EndpotintUnFollowChannel                              = EndpointAPIV1 + "unfollowChannel"
	EndpotintSetPrioritizedPaymentMethod                  = EndpointAPIV1 + "setPrioritizedPaymentMethod"
	EndpotintAgreeToPolicy                                = EndpointAPIV1 + "agreeToPolicy"
	EndpotintRejectP2PRequestMoney                        = EndpointAPIV1 + "rejectP2PRequestMoney"
	EndpotintExecuteTopup                                 = EndpointAPIV1 + "executeTopup"
	EndpotintPreAuthDisplayInfo                           = EndpointAPIV1 + "getPreAuthDisplayInfo"
	EndpotintCancelP2PSendMoney                           = EndpointAPIV1 + "cancelP2PSendMoney"
	EndpotintLoginDevices                                 = EndpointAPIV1 + "getLoginDevices"
	EndpotintDeletePaymentMethod                          = EndpointAPIV1 + "deletePaymentMethod"
	EndpotintGiftCardInfo                                 = EndpointAPIV1 + "getGiftCardInfo"
	EndpotintRenewP2PSession                              = EndpointAPIV1 + "renewP2PSession"
	EndpotintSignOut                                      = EndpointAPIV1 + "signOut"
	EndpotintLinkYconnect                                 = EndpointAPIV1 + "linkYconnect"
	EndpotintPPStepPaymentHistory                         = EndpointAPIV1 + "getPPStepPaymentHistory"
	EndpotintAcceptP2PSendMoney                           = EndpointAPIV1 + "acceptP2PSendMoney"
	EndpotintPayLaterInfo                                 = EndpointAPIV1 + "getPayLaterInfo"
	EndpotintChangeServiceLinkState                       = EndpointAPIV1 + "changeServiceLinkState"
	EndpotintGVList                                       = EndpointAPIV1 + "getGVList"
	EndpotintSendOtpForUserDefinedLimits                  = EndpointAPIV1 + "sendOtpForUserDefinedLimits"
	EndpotintPaymentCompletion                            = EndpointAPIV1 + "getPaymentCompletion"
	EndpotintP2PMessage                                   = EndpointAPIV1 + "getP2PMessage"

	EndpotintProfileDisplayInfo                     = EndpointAPIV2 + "getProfileDisplayInfo"
	EndpotintWalletDisplayInfo                      = EndpointAPIV2 + "getWalletDisplayInfo"
	EndpotintPrioritizedPaymentMethodsConfiguration = EndpointAPIV2 + "getPrioritizedPaymentMethodsConfiguration"
	EndpotintAutoTopupConfiguration                 = EndpointAPIV2 + "getAutoTopupConfiguration"
	EndpotintPayoutInfo                             = EndpointAPIV2 + "getPayoutInfo"
	EndpotintPaymentMethodList                      = EndpointAPIV2 + "getPaymentMethodList"
	EndpotintBarcodeInfo                            = EndpointAPIV2 + "getBarcodeInfo"
	EndpotintRejectP2PSendMoneyLink                 = EndpointAPIV2 + "rejectP2PSendMoneyLink"
	EndpotintUpdateAutoTopupConfiguration           = EndpointAPIV2 + "updateAutoTopupConfiguration"
	EndpotintV2PayoutDisplayInfo                    = EndpointAPIV2 + "updateAutoTopupConfiguration"
	EndpotintEvxecuteP2PRequestMoney                = EndpointAPIV2 + "executeP2PRequestMoney"
	EndpotintHomeDisplayInfo                        = EndpointAPIV2 + "getHomeDisplayInfo"
	EndpotintAcceptP2PSendMoneyLink                 = EndpointAPIV2 + "acceptP2PSendMoneyLink"
	EndpotintExecuteP2PSendMoneyLink                = EndpointAPIV2 + "executeP2PSendMoneyLink"
	EndpotintPay2BalanceHistory                     = EndpointAPIV2 + "getPay2BalanceHistory"
	EndpotintP2PLinkInfo                            = EndpointAPIV2 + "getP2PLinkInfo"
	EndpotintTopupDisplayInfo                       = EndpointAPIV2 + "getTopupDisplayInfo"

	EndpointOauth2      = EndpointAPIV2 + "/oauth2/"
	EndpointOauth2Par   = EndpointOauth2 + "par"
	EndpointOauth2Token = EndpointOauth2 + "token"

	EndpointCreatePaymentOneTimeCode        = EndpointAPIV3 + "createPaymentOneTimeCode"
	EndpointCreatePaymentOneTimeCodeForHome = EndpointAPIV3 + "createPaymentOneTimeCodeForHome"
	EndpointCashBackResult                  = EndpointAPIV3 + "getCashBackResult"
	EndpointPaymentHistory                  = EndpointAPIV3 + "getCashBackResult"
)
