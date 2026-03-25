import { Locator, Page, expect } from "@playwright/test";
import { BasePage } from "../basePage";

export class StatsPage extends BasePage {
    protected pageName = "Страница статистики";

    readonly statsLink: Locator;
    readonly refreshButton: Locator;
    readonly toggleButton: Locator;

    readonly timeLabel: Locator;
    readonly timeValue: Locator;

    constructor(page: Page) {
        super(page);

        this.statsLink = page.locator("a[href='/stats']");
        this.refreshButton = page.locator("button[aria-label='Обновить сейчас']");
        this.toggleButton = page.locator("button[aria-label*='автообновление']");

        this.timeLabel = page.locator("._timeLabel_ir5wu_106");
        this.timeValue = page.locator("._timeValue_ir5wu_112");
    }

    protected root(): Locator {
        return this.refreshButton;
    }

    async open() {
        await this.page.goto("/");
        await this.statsLink.click();
        await this.page.waitForLoadState("networkidle");
        await this.waitForOpen();
    }

    async getTimerText(): Promise<string> {
        await expect(this.timeValue).toBeVisible();
        return (await this.timeValue.textContent()) ?? "";
    }

    async refreshAndWaitForReset(previous: string) {
        await this.refreshButton.click();

        await expect.poll(async () => {
            const current = await this.timeValue.textContent();
            return current;
        }).not.toBe(previous);
    }

    async stopTimer() {
        await this.toggleButton.click();

        // таймер должен исчезнуть
        await expect(this.timeLabel).toHaveCount(0);
        await expect(this.timeValue).toHaveCount(0);

        // кнопка изменилась на "включить"
        await expect(this.toggleButton).toHaveAttribute(
            "aria-label",
            /Включить|Запустить/
        );
    }

    async startTimer() {
        await this.toggleButton.click();

        // таймер снова появился
        await expect(this.timeValue).toBeVisible();

        // кнопка снова "пауза"
        await expect(this.toggleButton).toHaveAttribute(
            "aria-label",
            /Отключить/
        );
    }

    parseTimeToSeconds(time: string): number {
        const [m, s] = time.split(":").map(Number);
        return m * 60 + s;
    }
}
