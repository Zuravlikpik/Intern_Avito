import { Page, Locator, expect } from "@playwright/test";

export abstract class BasePage {
    protected constructor(protected page: Page) {}

    protected abstract root(): Locator;

    async waitForOpen() {
        await expect(this.root()).toBeVisible();
    }
}
