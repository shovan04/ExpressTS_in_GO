package shared

func GetHttpResponseCodeContent() []byte {
	return []byte(`export default class HttpResponseCode {
    // 1xx: Informational
    static CONTINUE: number = 100;
    static SWITCHING_PROTOCOLS: number = 101;
    static PROCESSING: number = 102;
    static EARLY_HINTS: number = 103;

    // 2xx: Success
    static OK: number = 200;
    static CREATED: number = 201;
    static ACCEPTED: number = 202;
    static NON_AUTHORITATIVE_INFORMATION: number = 203;
    static NO_CONTENT: number = 204;
    static RESET_CONTENT: number = 205;
    static PARTIAL_CONTENT: number = 206;
    static MULTI_STATUS: number = 207;
    static ALREADY_REPORTED: number = 208;
    static IM_USED: number = 226;

    // 3xx: Redirection
    static MULTIPLE_CHOICES: number = 300;
    static MOVED_PERMANENTLY: number = 301;
    static FOUND: number = 302;
    static SEE_OTHER: number = 303;
    static NOT_MODIFIED: number = 304;
    static USE_PROXY: number = 305;
    static SWITCH_PROXY: number = 306; // Deprecated
    static TEMPORARY_REDIRECT: number = 307;
    static PERMANENT_REDIRECT: number = 308;

    // 4xx: Client Errors
    static BAD_REQUEST: number = 400;
    static UNAUTHORIZED: number = 401;
    static PAYMENT_REQUIRED: number = 402;
    static FORBIDDEN: number = 403;
    static NOTFOUND: number = 404; // keep your original name
    static METHOD_NOT_ALLOWED: number = 405;
    static NOT_ACCEPTABLE: number = 406;
    static PROXY_AUTHENTICATION_REQUIRED: number = 407;
    static REQUEST_TIMEOUT: number = 408;
    static CONFLICT: number = 409;
    static GONE: number = 410;
    static LENGTH_REQUIRED: number = 411;
    static PRECONDITION_FAILED: number = 412;
    static PAYLOAD_TOO_LARGE: number = 413;
    static URI_TOO_LONG: number = 414;
    static UNSUPPORTED_MEDIA_TYPE: number = 415;
    static RANGE_NOT_SATISFIABLE: number = 416;
    static EXPECTATION_FAILED: number = 417;
    static IM_A_TEAPOT: number = 418; // RFC 2324 joke
    static MISDIRECTED_REQUEST: number = 421;
    static UNPROCESSABLE_ENTITY: number = 422;
    static LOCKED: number = 423;
    static FAILED_DEPENDENCY: number = 424;
    static TOO_EARLY: number = 425;
    static UPGRADE_REQUIRED: number = 426;
    static PRECONDITION_REQUIRED: number = 428;
    static TOO_MANY_REQUESTS: number = 429;
    static REQUEST_HEADER_FIELDS_TOO_LARGE: number = 431;
    static UNAVAILABLE_FOR_LEGAL_REASONS: number = 451;

    // 5xx: Server Errors
    static INTERNAL_SERVER_ERROR: number = 500;
    static NOT_IMPLEMENTED: number = 501;
    static BAD_GATEWAY: number = 502;
    static SERVICE_UNAVAILABLE: number = 503;
    static GATEWAY_TIMEOUT: number = 504;
    static HTTP_VERSION_NOT_SUPPORTED: number = 505;
    static VARIANT_ALSO_NEGOTIATES: number = 506;
    static INSUFFICIENT_STORAGE: number = 507;
    static LOOP_DETECTED: number = 508;
    static NOT_EXTENDED: number = 510;
    static NETWORK_AUTHENTICATION_REQUIRED: number = 511;
}
`)}