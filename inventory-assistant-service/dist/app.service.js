"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.AppService = void 0;
const common_1 = require("@nestjs/common");
const path_1 = require("path");
const index_helper_1 = require("./tools/index.helper");
let AppService = class AppService {
    getHello() {
        return 'Hello World!';
    }
    getMatchInfo() {
        const rootPath = '/Users/apple/fzhange/inventory-assistant/inventory-assistant-service';
        const rawInventoryData = (0, index_helper_1.readExcelFile)((0, path_1.join)(rootPath, 'src/assets/inventory.xlsx'));
        const rawOrderData = (0, index_helper_1.readExcelFile)((0, path_1.join)(rootPath, 'src/assets/order.xlsx'));
        const cookedOrderData = (0, index_helper_1.cookOrderData)(rawOrderData);
        const cookedInventoryData = (0, index_helper_1.cookInventoryData)(rawInventoryData);
        return (0, index_helper_1.matchOrderAndInventory)(cookedOrderData, cookedInventoryData);
    }
};
exports.AppService = AppService;
exports.AppService = AppService = __decorate([
    (0, common_1.Injectable)()
], AppService);
//# sourceMappingURL=app.service.js.map