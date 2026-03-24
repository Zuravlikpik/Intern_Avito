import { test, expect } from "@playwright/test";
import { ListPage } from "../../pages/listPage/listPage";

test.describe("Фильтр по категории", () => {

    test("Все категории (по умолчанию)", async ({ page }) => {
        const listPage = new ListPage(page);

        await listPage.open();

        // убеждаемся что выбрано "Все категории"
        await listPage.selectCategory("");

        let hasNext = true;
        let pageCount = 0;

        while (hasNext && pageCount < 3) {
            const categories = await listPage.getCategories();

            expect(categories.length).toBeGreaterThan(0);

            // проверяем что есть разные категории
            await listPage.assertMultipleCategories(categories);

            hasNext = await listPage.goToNextPageIfExists();
            pageCount++;
        }
    });

    const categoryMap = [
        { value: "0", name: "Электроника" },
        { value: "1", name: "Недвижимость" },
        { value: "2", name: "Транспорт" },
        { value: "3", name: "Работа" },
        { value: "4", name: "Услуги" },
        { value: "5", name: "Животные" },
        { value: "6", name: "Мода" },
        { value: "7", name: "Детское" },
    ];

    for (const category of categoryMap) {
        test(`Категория: ${category.name}`, async ({ page }) => {
            const listPage = new ListPage(page);

            await listPage.open();

            await listPage.selectCategory(category.value);

            let hasNext = true;
            let pageCount = 0;

            while (hasNext && pageCount < 3) {
                const categories = await listPage.getCategories();

                expect(categories.length).toBeGreaterThan(0);

                // проверяем что ВСЕ карточки нужной категории
                await listPage.assertSingleCategory(categories, category.name);

                hasNext = await listPage.goToNextPageIfExists();
                pageCount++;
            }
        });
    }
});