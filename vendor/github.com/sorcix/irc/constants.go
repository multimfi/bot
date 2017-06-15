// Copyright 2014 Vic Demuzere
//
// Use of this source code is governed by the MIT license.

package irc

// Various prefixes extracted from RFC1459.
const (
	Channel     = '#' // Normal channel
	Distributed = '&' // Distributed channel

	Owner        = '~' // Channel owner +q (non-standard)
	Admin        = '&' // Channel admin +a (non-standard)
	Operator     = '@' // Channel operator +o
	HalfOperator = '%' // Channel half operator +h (non-standard)
	Voice        = '+' // User has voice +v
)

// User modes as defined by RFC1459 section 4.2.3.2.
const (
	UserModeInvisible     = 'i' // User is invisible
	UserModeServerNotices = 's' // User wants to receive server notices
	UserModeWallops       = 'w' // User wants to receive Wallops
	UserModeOperator      = 'o' // Server operator
)

// Channel modes as defined by RFC1459 section 4.2.3.1
const (
	ModeOperator   = 'o' // Operator privileges
	ModeVoice      = 'v' // Ability to speak on a moderated channel
	ModePrivate    = 'p' // Private channel
	ModeSecret     = 's' // Secret channel
	ModeInviteOnly = 'i' // Users can't join without invite
	ModeTopic      = 't' // Topic can only be set by an operator
	ModeModerated  = 'm' // Only voiced users and operators can talk
	ModeLimit      = 'l' // User limit
	ModeKey        = 'k' // Channel password

	ModeOwner        = 'q' // Owner privileges (non-standard)
	ModeAdmin        = 'a' // Admin privileges (non-standard)
	ModeHalfOperator = 'h' // Half-operator privileges (non-standard)
)

// IRC commands extracted from RFC2812 section 3 and RFC2813 section 4.
const (
	CmdPass     = "PASS"
	CmdNick     = "NICK"
	CmdUser     = "USER"
	CmdOper     = "OPER"
	CmdMode     = "MODE"
	CmdService  = "SERVICE"
	CmdQuit     = "QUIT"
	CmdSQuit    = "SQUIT"
	CmdJoin     = "JOIN"
	CmdPart     = "PART"
	CmdTopic    = "TOPIC"
	CmdNames    = "NAMES"
	CmdList     = "LIST"
	CmdInvite   = "INVITE"
	CmdKick     = "KICK"
	CmdPrivMsg  = "PRIVMSG"
	CmdNotice   = "NOTICE"
	CmdMotd     = "MOTD"
	CmdLUsers   = "LUSERS"
	CmdVersion  = "VERSION"
	CmdStats    = "STATS"
	CmdLinks    = "LINKS"
	CmdTime     = "TIME"
	CmdConnect  = "CONNECT"
	CmdTrace    = "TRACE"
	CmdAdmin    = "ADMIN"
	CmdInfo     = "INFO"
	CmdServList = "SERVLIST"
	CmdSQuery   = "SQUERY"
	CmdWho      = "WHO"
	CmdWhois    = "WHOIS"
	CmdWhoWas   = "WHOWAS"
	CmdKill     = "KILL"
	CmdPing     = "PING"
	CmdPong     = "PONG"
	CmdError    = "ERROR"
	CmdAway     = "AWAY"
	CmdRehash   = "REHASH"
	CmdDie      = "DIE"
	CmdRestart  = "RESTART"
	CmdSummon   = "SUMMON"
	CmdUsers    = "USERS"
	CmdWallops  = "WALLOPS"
	CmdUserHost = "USERHOST"
	CmdIsOn     = "ISON"
	CmdServer   = "SERVER"
	CmdNJoin    = "NJOIN"
)

// Numeric IRC replies extracted from RFC2812 section 5.
const (
	ReplyWelcome         = "001" // RPL_WELCOME
	ReplyYourHost        = "002" // RPL_YOURHOST
	ReplyCreated         = "003" // RPL_CREATED
	ReplyMyInfo          = "004" // RPL_MYINFO
	ReplyBounce          = "005" // RPL_BOUNCE
	ReplyISupport        = "005" // RPL_ISUPPORT
	ReplyUserHost        = "302" // RPL_USERHOST
	ReplyIsOn            = "303" // RPL_ISON
	ReplyAway            = "301" // RPL_AWAY
	ReplyUnAway          = "305" // RPL_UNAWAY
	ReplyNowAway         = "306" // RPL_NOWAWAY
	ReplyWhoisUser       = "311" // RPL_WHOISUSER
	ReplyWhoisServer     = "312" // RPL_WHOISSERVER
	ReplyWhoisOperator   = "313" // RPL_WHOISOPERATOR
	ReplyWhoisIdle       = "317" // RPL_WHOISIDLE
	ReplyEndofWhois      = "318" // RPL_ENDOFWHOIS
	ReplyWhoisChannels   = "319" // RPL_WHOISCHANNELS
	ReplyWhowasUser      = "314" // RPL_WHOWASUSER
	ReplyEndofWhowas     = "369" // RPL_ENDOFWHOWAS
	ReplyListStart       = "321" // RPL_LISTSTART
	ReplyList            = "322" // RPL_LIST
	ReplyListEnd         = "323" // RPL_LISTEND
	ReplyUniqOpIs        = "325" // RPL_UNIQOPIS
	ReplyChannelModeIs   = "324" // RPL_CHANNELMODEIS
	ReplyNoTopic         = "331" // RPL_NOTOPIC
	ReplyTopic           = "332" // RPL_TOPIC
	ReplyInviting        = "341" // RPL_INVITING
	ReplySummoning       = "342" // RPL_SUMMONING
	ReplyInviteList      = "346" // RPL_INVITELIST
	ReplyEndofInviteList = "347" // RPL_ENDOFINVITELIST
	ReplyExceptList      = "348" // RPL_EXCEPTLIST
	ReplyEndofExceptList = "349" // RPL_ENDOFEXCEPTLIST
	ReplyVersion         = "351" // RPL_VERSION
	ReplyWhoReply        = "352" // RPL_WHOREPLY
	ReplyEndofWho        = "315" // RPL_ENDOFWHO
	ReplyNamReply        = "353" // RPL_NAMREPLY
	ReplyEndofNames      = "366" // RPL_ENDOFNAMES
	ReplyLinks           = "364" // RPL_LINKS
	ReplyEndofLinks      = "365" // RPL_ENDOFLINKS
	ReplyBanList         = "367" // RPL_BANLIST
	ReplyEndofBanList    = "368" // RPL_ENDOFBANLIST
	ReplyInfo            = "371" // RPL_INFO
	ReplyEndofInfo       = "374" // RPL_ENDOFINFO
	ReplyMotdStart       = "375" // RPL_MOTDSTART
	ReplyMotd            = "372" // RPL_MOTD
	ReplyEndofMotd       = "376" // RPL_ENDOFMOTD
	ReplyYoureOper       = "381" // RPL_YOUREOPER
	ReplyRehashing       = "382" // RPL_REHASHING
	ReplyYoureService    = "383" // RPL_YOURESERVICE
	ReplyTime            = "391" // RPL_TIME
	ReplyUsersStart      = "392" // RPL_USERSSTART
	ReplyUsers           = "393" // RPL_USERS
	ReplyEndofUsers      = "394" // RPL_ENDOFUSERS
	ReplyNoUsers         = "395" // RPL_NOUSERS
	ReplyTraceLink       = "200" // RPL_TRACELINK
	ReplyTraceConnecting = "201" // RPL_TRACECONNECTING
	ReplyTraceHandshake  = "202" // RPL_TRACEHANDSHAKE
	ReplyTraceUnknown    = "203" // RPL_TRACEUNKNOWN
	ReplyTraceOperator   = "204" // RPL_TRACEOPERATOR
	ReplyTraceUser       = "205" // RPL_TRACEUSER
	ReplyTraceServer     = "206" // RPL_TRACESERVER
	ReplyTraceService    = "207" // RPL_TRACESERVICE
	ReplyTraceNewType    = "208" // RPL_TRACENEWTYPE
	ReplyTraceClass      = "209" // RPL_TRACECLASS
	ReplyTraceReconnect  = "210" // RPL_TRACERECONNECT
	ReplyTraceLog        = "261" // RPL_TRACELOG
	ReplyTraceEnd        = "262" // RPL_TRACEEND
	ReplyStatsLinkInfo   = "211" // RPL_STATSLINKINFO
	ReplyStatsCommands   = "212" // RPL_STATSCOMMANDS
	ReplyEndofStats      = "219" // RPL_ENDOFSTATS
	ReplyStatsUptime     = "242" // RPL_STATSUPTIME
	ReplyStatsOline      = "243" // RPL_STATSOLINE
	ReplyUModeIs         = "221" // RPL_UMODEIS
	ReplyServList        = "234" // RPL_SERVLIST
	ReplyServListEnd     = "235" // RPL_SERVLISTEND
	ReplyLUserClient     = "251" // RPL_LUSERCLIENT
	ReplyLUserOp         = "252" // RPL_LUSEROP
	ReplyLUserUnknown    = "253" // RPL_LUSERUNKNOWN
	ReplyLUserChannels   = "254" // RPL_LUSERCHANNELS
	ReplyLUserMe         = "255" // RPL_LUSERME
	ReplyAdminMe         = "256" // RPL_ADMINME
	ReplyAdminLoc1       = "257" // RPL_ADMINLOC1
	ReplyAdminLoc2       = "258" // RPL_ADMINLOC2
	ReplyAdminEmail      = "259" // RPL_ADMINEMAIL
	ReplyTryAgain        = "263" // RPL_TRYAGAIN

	ErrorNoSuchNick        = "401" // ERR_NOSUCHNICK
	ErrorNoSuchServer      = "402" // ERR_NOSUCHSERVER
	ErrorNoSuchChannel     = "403" // ERR_NOSUCHCHANNEL
	ErrorCannotSendtoChan  = "404" // ERR_CANNOTSENDTOCHAN
	ErrorTooManyChannels   = "405" // ERR_TOOMANYCHANNELS
	ErrorWasNoSuchNick     = "406" // ERR_WASNOSUCHNICK
	ErrorTooManyTargets    = "407" // ERR_TOOMANYTARGETS
	ErrorNoSuchService     = "408" // ERR_NOSUCHSERVICE
	ErrorNoOrigin          = "409" // ERR_NOORIGIN
	ErrorNoRecipient       = "411" // ERR_NORECIPIENT
	ErrorNoTextToSend      = "412" // ERR_NOTEXTTOSEND
	ErrorNoTopLevel        = "413" // ERR_NOTOPLEVEL
	ErrorWildTopLevel      = "414" // ERR_WILDTOPLEVEL
	ErrorBadMask           = "415" // ERR_BADMASK
	ErrorUnknownCommand    = "421" // ERR_UNKNOWNCOMMAND
	ErrorNoMOTD            = "422" // ERR_NOMOTD
	ErrorNoAdminInfo       = "423" // ERR_NOADMININFO
	ErrorFileError         = "424" // ERR_FILEERROR
	ErrorNoNicknameGiven   = "431" // ERR_NONICKNAMEGIVEN
	ErrorErroneusNickname  = "432" // ERR_ERRONEUSNICKNAME
	ErrorNicknameInUse     = "433" // ERR_NICKNAMEINUSE
	ErrorNickCollision     = "436" // ERR_NICKCOLLISION
	ErrorUnavailResource   = "437" // ERR_UNAVAILRESOURCE
	ErrorUserNotInChannel  = "441" // ERR_USERNOTINCHANNEL
	ErrorNotOnChannel      = "442" // ERR_NOTONCHANNEL
	ErrorUserOnChannel     = "443" // ERR_USERONCHANNEL
	ErrorNoLogin           = "444" // ERR_NOLOGIN
	ErrorSummonDisabled    = "445" // ERR_SUMMONDISABLED
	ErrorUsersDisabled     = "446" // ERR_USERSDISABLED
	ErrorNotRegistered     = "451" // ERR_NOTREGISTERED
	ErrorNeedMoreParams    = "461" // ERR_NEEDMOREPARAMS
	ErrorAlreadyRegistred  = "462" // ERR_ALREADYREGISTRED
	ErrorNoPermForHost     = "463" // ERR_NOPERMFORHOST
	ErrorPasswdMismatch    = "464" // ERR_PASSWDMISMATCH
	ErrorYoureBannedCreep  = "465" // ERR_YOUREBANNEDCREEP
	ErrorYouWillBeBanned   = "466" // ERR_YOUWILLBEBANNED
	ErrorKeyset            = "467" // ERR_KEYSET
	ErrorChannelIsFull     = "471" // ERR_CHANNELISFULL
	ErrorUnknownMode       = "472" // ERR_UNKNOWNMODE
	ErrorInviteOnlyChan    = "473" // ERR_INVITEONLYCHAN
	ErrorBannedFromChan    = "474" // ERR_BANNEDFROMCHAN
	ErrorBadChannelKey     = "475" // ERR_BADCHANNELKEY
	ErrorBadChanMask       = "476" // ERR_BADCHANMASK
	ErrorNoChanModes       = "477" // ERR_NOCHANMODES
	ErrorBanListFull       = "478" // ERR_BANLISTFULL
	ErrorNoPrivileges      = "481" // ERR_NOPRIVILEGES
	ErrorChanoPrivsNeeded  = "482" // ERR_CHANOPRIVSNEEDED
	ErrorCantKillServer    = "483" // ERR_CANTKILLSERVER
	ErrorRestricted        = "484" // ERR_RESTRICTED
	ErrorUniqOpPrivsNeeded = "485" // ERR_UNIQOPPRIVSNEEDED
	ErrorNoOperHost        = "491" // ERR_NOOPERHOST
	ErrorUModeUnknownFlag  = "501" // ERR_UMODEUNKNOWNFLAG
	ErrorUsersDontMatch    = "502" // ERR_USERSDONTMATCH
)

// IRC commands extracted from the IRCv3 spec at http://www.ircv3.org/.
const (
	Cap      = "CAP"
	CapLs    = "LS"    // CAP_LS    Subcommand (param)
	CapList  = "LIST"  // CAP_LIST  Subcommand (param)
	CapReq   = "REQ"   // CAP_REQ   Subcommand (param)
	CapAck   = "ACK"   // CAP_ACK   Subcommand (param)
	CapNak   = "NAK"   // CAP_NAK   Subcommand (param)
	CapClear = "CLEAR" // CAP_CLEAR Subcommand (param)
	CapEnd   = "END"   // CAP_END   Subcommand (param)

	Authenticate = "AUTHENTICATE"
)

// Numeric IRC replies extracted from the IRCv3 spec.
const (
	ReplyLoggedIn    = "900" // RPL_LOGGEDIN
	ReplyLoggedOut   = "901" // RPL_LOGGEDOUT
	ReplyNickLocked  = "902" // RPL_NICKLOCKED
	ReplySASLSuccess = "903" // RPL_SASLSUCCESS

	ErrorSASLFail    = "904" // ERR_SASLFAIL
	ErrorSASLToolong = "905" // ERR_SASLTOOLONG
	ErrorSASLAborted = "906" // ERR_SASLABORTED
	ErrorSASLAlready = "907" // ERR_SASLALREADY

	ReplySASLMechs = "908" // RPL_SASLMECHS
)

// RFC2812, section 5.3
const (
	ReplyStatsCline    = "213" // RPL_STATSCLINE
	ReplyStatsNline    = "214" // RPL_STATSNLINE
	ReplyStatsIline    = "215" // RPL_STATSILINE
	ReplyStatsKline    = "216" // RPL_STATSKLINE
	ReplyStatsQline    = "217" // RPL_STATSQLINE
	ReplyStatsYline    = "218" // RPL_STATSYLINE
	ReplyServiceInfo   = "231" // RPL_SERVICEINFO
	ReplyEndofServices = "232" // RPL_ENDOFSERVICES
	ReplyService       = "233" // RPL_SERVICE
	ReplyStatsVline    = "240" // RPL_STATSVLINE
	ReplyStatsLline    = "241" // RPL_STATSLLINE
	ReplyStatsHline    = "244" // RPL_STATSHLINE
	ReplyStatsSline    = "245" // RPL_STATSSLINE
	ReplyStatsPing     = "246" // RPL_STATSPING
	ReplyStatsBline    = "247" // RPL_STATSBLINE
	ReplyStatsDline    = "250" // RPL_STATSDLINE
	ReplyNone          = "300" // RPL_NONE
	ReplyWhoisChanOp   = "316" // RPL_WHOISCHANOP
	ReplyKillDone      = "361" // RPL_KILLDONE
	ReplyClosing       = "362" // RPL_CLOSING
	ReplyCloseEnd      = "363" // RPL_CLOSEEND
	ReplyInfoStart     = "373" // RPL_INFOSTART
	ReplyMyPortIs      = "384" // RPL_MYPORTIS

	ErrorNoServiceHost = "492" // ERR_NOSERVICEHOST
)

// Other constants
const (
	ErrorTooManyMatches = "416" // ERR_TOOMANYMATCHES Used on IRCNet.

	ReplyTopicWhoTime = "333" // RPL_TOPICWHOTIME From ircu, in use on Freenode.
	ReplyLocalUsers   = "265" // RPL_LOCALUSERS From aircd, Hybrid, Hybrid, Bahamut, in use on Freenode.
	ReplyGlobalUsers  = "266" // RPL_GLOBALUSERS From aircd, Hybrid, Hybrid, Bahamut, in use on Freenode.
)
