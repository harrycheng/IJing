import { getDivinatory } from '~/server/utils/services/divinatory';

export default defineEventHandler(async (event) => {
    const today = new Date().toISOString().split('T')[0];
    const cookies = parseCookies(event);

    if (cookies && cookies['divination-name']) {
        try {
            const cachedDivinatory = JSON.parse(Buffer.from(cookies['divination-name'], 'base64').toString('utf-8'));

            if (cachedDivinatory.date === today) {
                // Cookie is valid for today. Return cached name and fetch fresh details.
                const fullData = await getDivinatory(cachedDivinatory.name);
                return {
                    success: true,
                    data: {
                        divinatory: cachedDivinatory.name,
                        divinatoryDetail: fullData.divinatoryDetail
                    }
                };
            }
        } catch (error) {
            console.error('Failed to parse cookie data:', error);
        }
    }

    // If no valid cookie, generate new divination
    const newData = await getDivinatory();
    
    // Create a smaller cookie with only the name and date
    const cookieValue = {
        date: today,
        name: newData.divinatory 
    };

    const encodedValue = Buffer.from(JSON.stringify(cookieValue)).toString('base64');

    // Use a different cookie name to avoid conflicts with old, larger cookies
    setCookie(event, 'divination-name', encodedValue, {
        maxAge: 60 * 60 * 24, // 1 day
        sameSite: 'lax',
        path: '/',
    });

    return {
        success: true,
        data: {
            divinatory: newData.divinatory,
            divinatoryDetail: newData.divinatoryDetail
        }
    };
});
