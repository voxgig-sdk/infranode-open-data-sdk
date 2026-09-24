"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.InfranodeOpenDataError = void 0;
class InfranodeOpenDataError extends Error {
    isInfranodeOpenDataError = true;
    sdk = 'InfranodeOpenData';
    code;
    ctx;
    status = -1;
    // `err.notFound` rather than a magic number at every call site.
    get notFound() { return 404 === this.status; }
    constructor(code, msg, ctx) {
        super(msg);
        this.code = code;
        this.ctx = ctx;
    }
}
exports.InfranodeOpenDataError = InfranodeOpenDataError;
//# sourceMappingURL=InfranodeOpenDataError.js.map