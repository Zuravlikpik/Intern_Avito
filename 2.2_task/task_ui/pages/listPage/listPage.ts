import { Locator, Page, expect } from "@playwright/test";
import { BasePage } from "../basePage";

export class ListPage extends BasePage {
    protected pageName = "Страница объявлений";

    readonly sortSelect: Locator;
    readonly orderSelect: Locator;
    readonly priceList: Locator;
    readonly nextPageButton: Locator;
    readonly categorySelect: Locator;
    readonly categoryList: Locator;
    readonly urgentToggle: Locator;
    readonly prioritySelect: Locator;
    readonly urgentBadges: Locator;
    

    constructor(page: Page) {
        super(page);

        this.sortSelect = page.locator("._filters__select_1iunh_21").first();
        this.orderSelect = page.locator("._filters__select_1iunh_21").nth(1);
        this.priceList = page.locator("._card__price_15fhn_241");
        this.nextPageButton = page.locator("button[aria-label='Следующая страница']");
        this.categorySelect = page.locator("select").nth(2);
        this.categoryList = page.locator("._card__category_15fhn_259");
        this.urgentToggle = page.locator("label:has-text('Только срочные')");
        this.prioritySelect = page.locator("select").nth(3);
        this.urgentBadges = page.locator("text=Срочно");
    }

    protected root(): Locator {
        return this.page.locator("text=Сортировать по");
    }

    async open() {
        await this.page.goto("/");
        await this.page.waitForLoadState("networkidle");

        await this.page.locator("text=Сортировать по").waitFor();
    }

    async sortByPriceDesc() {
        await this.sortSelect.selectOption("price");
        await this.orderSelect.selectOption("desc");
        await this.page.waitForTimeout(500);
        await this.page.waitForLoadState("networkidle");
    }

    async getPrices(): Promise<number[]> {
        const texts = await this.priceList.allTextContents();

        return texts.map((text) =>
            Number(text.replace(/\s|₽/g, ""))
        );
    }

    async assertSortedDesc(prices: number[]) {
        for (let i = 0; i < prices.length - 1; i++) {
            expect(prices[i]).toBeGreaterThanOrEqual(prices[i + 1]);
        }
    }

    async goToNextPageIfExists(): Promise<boolean> {
        const button = this.nextPageButton;

        if (await button.count() === 0) {
            return false;
        }

        await this.page.waitForTimeout(500);

        if (!(await button.isEnabled())) {
            return false;
        }

        const currentFirstItem =
            (await this.priceList.first().textContent()) ?? "";

        await button.click();

        await this.page.waitForFunction(
            ([selector, prev]) => {
                const el = document.querySelector(selector);
                return el && el.textContent !== prev;
            },
            ["._card__price_15fhn_241", currentFirstItem]
        ).catch(() => {});

        await this.page.waitForLoadState("networkidle");

        return true;
    }

    async applyPriceRange(min: number, max: number) {
        const priceFromInput = this.page.locator("input[placeholder='От']");
        const priceToInput = this.page.locator("input[placeholder='До']");

        await priceFromInput.fill(String(min));
        await priceToInput.fill(String(max));

        const applyButton = this.page.locator("button:text('Применить')");
        if (await applyButton.count() > 0) {
            await applyButton.click();
        }
        await this.page.waitForLoadState("networkidle");
        await this.page.waitForTimeout(500);
    }


    async applyPriceRangeFilter(from: number, to: number) {
        const fromInput = this.page.locator("input[placeholder='От']");
        const toInput = this.page.locator("input[placeholder='До']");

        await fromInput.fill(String(from));
        await toInput.fill(String(to));

        await this.page.waitForLoadState("networkidle");
        await this.page.waitForTimeout(500);
    }

    async selectCategory(value: string) {
        await this.categorySelect.selectOption(value);
        await this.page.waitForLoadState("networkidle");
        await this.page.waitForTimeout(500);
    }

    async getCategories(): Promise<string[]> {
        const texts = await this.categoryList.allTextContents();
        return texts.map(t => t.trim());
    }

    async assertMultipleCategories(categories: string[]) {
        const unique = new Set(categories);
        expect(unique.size).toBeGreaterThan(1);
    }

    async assertSingleCategory(categories: string[], expected: string) {
        for (const category of categories) {
            expect(category.toLowerCase()).toContain(expected.toLowerCase());
        }
    }

    async enableUrgentOnly() {
        await this.urgentToggle.click();

        const checkbox = this.page.locator("input._urgentToggle__input_h1vv9_14");
        await expect(checkbox).toBeChecked();

        await this.page.waitForLoadState("networkidle");
    }

    async assertPriorityUrgent() {
        await expect(this.prioritySelect).toHaveValue("urgent");
    }

    async getCardsCount(): Promise<number> {
        return await this.priceList.count();
    }

    async getUrgentCount(): Promise<number> {
        return await this.urgentBadges.count();
    }
}
